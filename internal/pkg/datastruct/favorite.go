package datastruct

const FavoriteTableName = "favorite"

type FavoriteItem struct {
	ID        int64 `db:"id"`
	UserID    int64 `db:"user_id"`
	ProductID int64 `db:"product_id"`
}

type FavoriteProduct struct {
	FavoriteID  int64
	Name        string
	Description string
	Url         string
	BrandID     int64
	CategoryID  int64
	Price       float64
	Color       string
}
