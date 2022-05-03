package datastruct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

const OrderTableName = "orders"

type Order struct {
	ID          int64       `db:"id"`
	UserID      int64       `db:"user_id"`
	OrderStatus OrderStatus `db:"order_status"`
	OrderDate   time.Time   `db:"order_date"`
	TotalPrice  float64     `db:"total_cost"`
	Items       orderItems  `db:"items"`
}

type orderItems []*FullCartItem

func (o orderItems) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *orderItems) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}

type UpdateOrder struct {
	ID          int64       `db:"id"`
	OrderStatus OrderStatus `db:"order_status"`
}

type OrderStatus string

const (
	OrderStatus_DONE     OrderStatus = "DONE"
	OrderStatus_CREATED  OrderStatus = "CREATED"
	OrderStatus_DECLINED OrderStatus = "DECLINED"
)
