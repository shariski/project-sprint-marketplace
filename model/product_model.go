package model

type ProductCreateModel struct {
	UserId         int      `json:"userId" validate:"required"`
	Name           string   `json:"name" validate:"required,min=5,max=60"`
	Price          float32  `json:"price" validate:"required,min=0"`
	ImageUrl       string   `json:"imageUrl" validate:"required,url"`
	Stock          int      `json:"stock" validate:"required,min=0"`
	Condition      string   `json:"condition" validate:"required,oneof=new second"`
	Tags           []string `json:"tags" validate:"required,min=0"`
	IsPurchaseable bool     `json:"isPurchaseable" validate:"required"`
}

type ProductUpdateModel struct {
	Id             int      `json:"id" validate:"required"`
	Name           string   `json:"name" validate:"required,min=5,max=60"`
	Price          float32  `json:"price" validate:"required,min=0"`
	ImageUrl       string   `json:"imageUrl" validate:"required,url"`
	Condition      string   `json:"condition" validate:"required,oneof=new second"`
	Tags           []string `json:"tags" validate:"required,min=0"`
	IsPurchaseable bool     `json:"isPurchaseable" validate:"required"`
}