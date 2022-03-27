package datastruct

const RoleTableName = "role"

type Role struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
