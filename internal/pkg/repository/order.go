package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type OrderQuery interface {
	Create(ctx context.Context, req datastruct.Order) (*datastruct.Order, error)
	Update(ctx context.Context, req datastruct.UpdateOrder) error
	Get(ctx context.Context, ID int64) (*datastruct.Order, error)
	List(ctx context.Context, userID int64) ([]*datastruct.Order, error)
}

type orderQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *orderQuery) Update(ctx context.Context, req datastruct.UpdateOrder) error {
	qb := q.builder.Update(datastruct.CartTableName).
		Where(squirrel.Eq{"id": req.ID}).
		Set("order_status", req.OrderStatus)

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

func (q *orderQuery) Create(ctx context.Context, req datastruct.Order) (*datastruct.Order, error) {
	qb := q.builder.Insert(datastruct.OrderTableName).
		Columns(
			"user_id",
			"order_status",
			"order_date",
			"total_cost",
			"items",
		).
		Values(
			req.UserID,
			req.OrderStatus,
			req.OrderDate,
			req.TotalPrice,
			req.Items,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var order datastruct.Order

	err = q.db.GetContext(ctx, &order, query, args...)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (q *orderQuery) Get(ctx context.Context, ID int64) (*datastruct.Order, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.OrderTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var order datastruct.Order

	err = q.db.GetContext(ctx, &order, query, args...)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (q *orderQuery) List(ctx context.Context, userID int64) ([]*datastruct.Order, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.OrderTableName).
		Where(squirrel.Eq{"user_id": userID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var orders []*datastruct.Order

	err = q.db.SelectContext(ctx, &orders, query, args...)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func NewOrderQuery(db *sqlx.DB) OrderQuery {
	return &orderQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
