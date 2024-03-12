package entity

type Product struct {
	Id             int     `sql:"id"`
	UserId         int     `sql:"user_id"`
	Name           string  `sql:"name"`
	Price          float32 `sql:"price"`
	ImageUrl       string  `sql:"image_url"`
	Stock          int     `sql:"stock"`
	Condition      string  `sql:"condition"`
	IsPurchaseable bool    `sql:"is_purchaseable"`
	CreatedAt      string  `sql:"created_at"`
	UpdatedAt      string  `sql:"updated_at"`
}
