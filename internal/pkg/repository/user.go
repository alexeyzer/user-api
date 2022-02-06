package repository

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type UserQuery interface {
	Create(ctx context.Context, req User) (*User, error)
	Get(ctx context.Context, email string) (*User, error)
	Exists(ctx context.Context, email string) (bool, error)
}

type userQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *userQuery) Create(ctx context.Context, req User) (*User, error) {
	qb := q.builder.Insert(UserTableName).
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

	var user User

	err = q.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (q *userQuery) Get(ctx context.Context, email string) (*User, error) {
	qb := q.builder.
		Select("*").
		From(UserTableName).
		Where(squirrel.Eq{"email": email})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var user User

	err = q.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (q *userQuery) Exists(ctx context.Context, username string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(UserTableName).
		Where(squirrel.Eq{"email": username})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var user User

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
