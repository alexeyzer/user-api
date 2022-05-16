package repository

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type RoleQuery interface {
	Create(ctx context.Context, req datastruct.Role) (*datastruct.Role, error)
	Update(ctx context.Context, req datastruct.Role) (*datastruct.Role, error)
	Get(ctx context.Context, ID int64) (*datastruct.Role, error)
	Exists(ctx context.Context, name string) (bool, error)
	List(ctx context.Context) ([]*datastruct.Role, error)
	Delete(ctx context.Context, ID int64) error
}

type roleQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *roleQuery) Update(ctx context.Context, req datastruct.Role) (*datastruct.Role, error) {
	qb := q.builder.Update(datastruct.RoleTableName).
		Set("name", req.Name).
		Set("description", req.Description).
		Where(squirrel.Eq{"id": req.ID}).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var role datastruct.Role

	err = q.db.GetContext(ctx, &role, query, args...)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (q *roleQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.RoleTableName).
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

func (q *roleQuery) Create(ctx context.Context, req datastruct.Role) (*datastruct.Role, error) {
	qb := q.builder.Insert(datastruct.RoleTableName).
		Columns(
			"name",
			"description",
		).
		Values(
			req.Name,
			req.Description,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var role datastruct.Role

	err = q.db.GetContext(ctx, &role, query, args...)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (q *roleQuery) Get(ctx context.Context, ID int64) (*datastruct.Role, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.RoleTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var role datastruct.Role

	err = q.db.GetContext(ctx, &role, query, args...)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (q *roleQuery) List(ctx context.Context) ([]*datastruct.Role, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.RoleTableName)
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var roles []*datastruct.Role

	err = q.db.SelectContext(ctx, &roles, query, args...)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (q *roleQuery) Exists(ctx context.Context, name string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.RoleTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var role datastruct.Role

	err = q.db.GetContext(ctx, &role, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewRoleQuery(db *sqlx.DB) RoleQuery {
	return &roleQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
