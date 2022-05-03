package datastruct

const UserTableName = "users"

type User struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Surname    string `db:"surname"`
	Patronymic string `db:"patronymic"`
	Password   []byte `db:"password"`
	Phone      string `db:"phone"`
	Email      string `db:"email"`
}

type UserWithRoles struct {
	ID    int64
	Email string
	Roles []string
}
