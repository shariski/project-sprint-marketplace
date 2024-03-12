package entity

type Tag struct {
	ProductId int    `sql:"product_id"`
	Name      string `sql:"name"`
	CreatedAt string `sql:"created_at"`
	UpdatedAt string `sql:"updated_at"`
}
