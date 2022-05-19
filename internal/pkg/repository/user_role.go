package repository

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type UserRoleQuery interface {
	Create(ctx context.Context, req datastruct.UserRole) (*datastruct.UserRole, error)
	Get(ctx context.Context, ID int64) (*datastruct.UserRole, error)
	Exists(ctx context.Context, userID, roleID int64) (bool, error)
	List(ctx context.Context, userID int64) ([]*datastruct.UserRoleWithName, error)
	Delete(ctx context.Context, ID int64) error
}

type userRoleQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *userRoleQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.UserRoleNameTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	_, err = q.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (q *userRoleQuery) Create(ctx context.Context, req datastruct.UserRole) (*datastruct.UserRole, error) {
	qb := q.builder.Insert(datastruct.UserRoleNameTableName).
		Columns(
			"user_id",
			"role_id",
		).
		Values(
			req.UserID,
			req.RoleID,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var userRoles datastruct.UserRole

	err = q.db.GetContext(ctx, &userRoles, query, args...)
	if err != nil {
		return nil, err
	}

	return &userRoles, nil
}

func (q *userRoleQuery) Get(ctx context.Context, ID int64) (*datastruct.UserRole, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.UserRoleNameTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var userRole datastruct.UserRole

	err = q.db.GetContext(ctx, &userRole, query, args...)
	if err != nil {
		return nil, err
	}

	return &userRole, nil
}

func (q *userRoleQuery) List(ctx context.Context, userID int64) ([]*datastruct.UserRoleWithName, error) {
	qb := q.builder.
		Select("urt.id, urt.user_id, urt.role_id, rt.name as role_name").
		From(datastruct.UserRoleNameTableName + " as urt").
		Where(squirrel.Eq{"user_id": userID}).
		LeftJoin(datastruct.RoleTableName + " as rt on rt.id = urt.role_id")

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var userRoles []*datastruct.UserRoleWithName

	err = q.db.SelectContext(ctx, &userRoles, query, args...)
	if err != nil {
		return nil, err
	}

	return userRoles, nil
}

func (q *userRoleQuery) Exists(ctx context.Context, userID, roleID int64) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.UserRoleNameTableName).
		Where(squirrel.And{squirrel.Eq{"user_id": userID}, squirrel.Eq{"role_id": roleID}})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var userRole datastruct.UserRole

	err = q.db.GetContext(ctx, &userRole, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewUserRolesQuery(db *sqlx.DB) UserRoleQuery {
	return &userRoleQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
