package middlewares

type Role struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
}

type UserRole struct {
	RoleId int `db:"role_id"`
}
