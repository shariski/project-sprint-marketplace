package model

type ProductCreateModel struct {
	UserId         int      `validate:"required"`
	Name           string   `validate:"required,min=5,max=60"`
	Price          float32  `validate:"required,min=0"`
	ImageUrl       string   `validate:"required,url"`
	Stock          int      `validate:"required,min=0"`
	Condition      string   `validate:"required,oneof=new second"`
	Tags           []string `validate:"required,min=0"`
	IsPurchaseable bool     `validate:"required"`
}
