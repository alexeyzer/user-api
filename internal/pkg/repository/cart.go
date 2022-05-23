package repository

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type CartQuery interface {
	Create(ctx context.Context, req datastruct.CartItem) (*datastruct.CartItem, error)
	Update(ctx context.Context, userID, finalProductID, quantity int64) (*datastruct.CartItem, error)
	Get(ctx context.Context, ID int64) (*datastruct.CartItem, error)
	Exists(ctx context.Context, userID, finalProductID int64) (bool, error)
	List(ctx context.Context, userID int64) ([]*datastruct.CartItem, error)
	Delete(ctx context.Context, ID int64) error
	DeleteByUserID(ctx context.Context, userID int64) error
}

type cartQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *cartQuery) DeleteByUserID(ctx context.Context, userID int64) error {
	qb := q.builder.
		Delete(datastruct.CartTableName).
		Where(squirrel.Eq{"user_id": userID})
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

func (q *cartQuery) Update(ctx context.Context, userID, finalProductID, quantity int64) (*datastruct.CartItem, error) {
	qb := q.builder.Update(datastruct.CartTableName).
		Where(squirrel.And{squirrel.Eq{"user_id": userID}, squirrel.Eq{"final_product_id": finalProductID}}).
		Set("quantity", quantity).
		Suffix("RETURNING *")

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var cartItem datastruct.CartItem

	err = q.db.GetContext(ctx, &cartItem, query, args...)
	if err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func (q *cartQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.CartTableName).
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

func (q *cartQuery) Create(ctx context.Context, req datastruct.CartItem) (*datastruct.CartItem, error) {
	qb := q.builder.Insert(datastruct.CartTableName).
		Columns(
			"user_id",
			"final_product_id",
			"quantity",
		).
		Values(
			req.UserID,
			req.FinalProductID,
			req.Quantity,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var cartItem datastruct.CartItem

	err = q.db.GetContext(ctx, &cartItem, query, args...)
	if err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func (q *cartQuery) Get(ctx context.Context, ID int64) (*datastruct.CartItem, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.CartTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var cartItem datastruct.CartItem

	err = q.db.GetContext(ctx, &cartItem, query, args...)
	if err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func (q *cartQuery) List(ctx context.Context, userID int64) ([]*datastruct.CartItem, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.CartTableName).
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("id asc")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var cartItems []*datastruct.CartItem

	err = q.db.SelectContext(ctx, &cartItems, query, args...)
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (q *cartQuery) Exists(ctx context.Context, userID, finalProductID int64) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.CartTableName).
		Where(squirrel.And{squirrel.Eq{"user_id": userID}, squirrel.Eq{"final_product_id": finalProductID}})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var cartItem datastruct.CartItem

	err = q.db.GetContext(ctx, &cartItem, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewCartQuery(db *sqlx.DB) CartQuery {
	return &cartQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
