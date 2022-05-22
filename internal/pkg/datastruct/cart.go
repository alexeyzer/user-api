package datastruct

const CartTableName = "cart"

type CartItem struct {
	ID             int64 `db:"id"`
	UserID         int64 `db:"user_id"`
	FinalProductID int64 `db:"final_product_id"`
	Quantity       int64 `db:"quantity"`
}

type FullCartItem struct {
	FullFinalProduct FullFinalProduct
	UserQuantity     int64
	ID               int64
}

type FullFinalProduct struct {
	ID           int64
	Amount       int64
	Sku          int64
	Name         string
	Description  string
	Url          string
	BrandName    string
	CategoryName string
	Price        float64
	Color        string
	Size         string
}
