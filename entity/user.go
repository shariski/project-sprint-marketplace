package entity

type User struct {
	Id        int    `sql:"id"`
	Username  string `sql:"username"`
	Name      string `sql:"name"`
	Password  string `sql:"password"`
	CreatedAt string `sql:"created_at"`
	UpdatedAt string `sql:"updated_at"`
}
