package repository

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserQuery interface {
	Create(ctx context.Context, req datastruct.User) (*datastruct.User, error)
	Get(ctx context.Context, email string) (*datastruct.User, error)
	GetByID(ctx context.Context, ID int64) (*datastruct.User, error)
	Exists(ctx context.Context, email string) (bool, error)
	List(ctx context.Context) ([]*datastruct.User, error)
}

type userQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *userQuery) List(ctx context.Context) ([]*datastruct.User, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.UserTableName)
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var users []*datastruct.User

	err = q.db.SelectContext(ctx, &users, query, args...)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (q *userQuery) Create(ctx context.Context, req datastruct.User) (*datastruct.User, error) {
	qb := q.builder.Insert(datastruct.UserTableName).
		Columns(
			"name",
			"surname",
			"patronymic",
			"password",
			"phone",
			"email",
		).
		Values(
			req.Name,
			req.Surname,
			req.Patronymic,
			req.Password,
			req.Phone,
			req.Email,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var user datastruct.User

	err = q.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (q *userQuery) GetByID(ctx context.Context, ID int64) (*datastruct.User, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.UserTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var user datastruct.User

	err = q.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (q *userQuery) Get(ctx context.Context, email string) (*datastruct.User, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.UserTableName).
		Where(squirrel.Eq{"email": email})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var user datastruct.User

	err = q.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.InvalidArgument, "User with email = %s doesn't exist", email)
		}
		return nil, err
	}

	return &user, nil
}

func (q *userQuery) Exists(ctx context.Context, username string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.UserTableName).
		Where(squirrel.Eq{"email": username})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var user datastruct.User

	err = q.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewUserQuery(db *sqlx.DB) UserQuery {
	return &userQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
