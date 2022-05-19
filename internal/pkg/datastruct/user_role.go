package datastruct

const UserRoleNameTableName = "user_role"

type UserRole struct {
	ID     int64 `db:"id"`
	UserID int64 `db:"user_id"`
	RoleID int64 `db:"role_id"`
}

type UserRoleWithName struct {
	ID       int64  `db:"id"`
	UserID   int64  `db:"user_id"`
	RoleID   int64  `db:"role_id"`
	RoleName string `db:"role_name"`
}
