package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type FavoriteQuery interface {
	Create(ctx context.Context, req datastruct.FavoriteItem) (*datastruct.FavoriteItem, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, userID int64) ([]*datastruct.FavoriteItem, error)
}

type favoriteQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *favoriteQuery) Create(ctx context.Context, req datastruct.FavoriteItem) (*datastruct.FavoriteItem, error) {
	qb := q.builder.Insert(datastruct.FavoriteTableName).
		Columns(
			"user_id",
			"product_id",
		).
		Values(
			req.UserID,
			req.ProductID,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var favoriteItem datastruct.FavoriteItem

	err = q.db.GetContext(ctx, &favoriteItem, query, args...)
	if err != nil {
		return nil, err
	}

	return &favoriteItem, nil
}

func (q *favoriteQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.FavoriteTableName).
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

func (q *favoriteQuery) List(ctx context.Context, userID int64) ([]*datastruct.FavoriteItem, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.FavoriteTableName).
		Where(squirrel.Eq{"user_id": userID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var favoriteItems []*datastruct.FavoriteItem

	err = q.db.SelectContext(ctx, &favoriteItems, query, args...)
	if err != nil {
		return nil, err
	}

	return favoriteItems, nil
}

func NewFavoriteQuery(db *sqlx.DB) FavoriteQuery {
	return &favoriteQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
