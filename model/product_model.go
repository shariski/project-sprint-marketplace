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
	UserId         int      `json:"userId" validate:"required"`
	Name           string   `json:"name" validate:"required,min=5,max=60"`
	Price          float32  `json:"price" validate:"required,min=0"`
	ImageUrl       string   `json:"imageUrl" validate:"required,url"`
	Condition      string   `json:"condition" validate:"required,oneof=new second"`
	Tags           []string `json:"tags" validate:"required,min=0"`
	IsPurchaseable bool     `json:"isPurchaseable" validate:"required"`
}

type ProductModel struct {
	Id             string   `json:"productId"`
	Name           string   `json:"name"`
	Price          float32  `json:"price"`
	ImageUrl       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
	PurchaseCount  int      `json:"purchaseCount"`
}

type GetProductModel struct {
	Product ProductModel `json:"product"`
	Seller  SellerModel  `json:"seller"`
}

type UpdateStockModel struct {
	Id     int `json:"id" validate:"required"`
	UserId int `json:"userId" validate:"required"`
	Stock  int `json:"stock" validate:"required,min=0"`
}

type ProductFilters struct {
	UserOnly       bool     `json:"userOnly" query:"userOnly"`
	UserId         int      `json:"userId"`
	Limit          int      `json:"limit" query:"limit"`
	Offset         int      `json:"offset" query:"offset"`
	Tags           []string `json:"tags" query:"tags"`
	Condition      string   `json:"condition" query:"condition"`
	ShowEmptyStock bool     `json:"showEmptyStock" query:"showEmptyStock"`
	MaxPrice       float64  `json:"maxPrice" query:"maxPrice"`
	MinPrice       float64  `json:"minPrice" query:"minPrice"`
	SortBy         string   `json:"sortBy" query:"sortBy"`
	OrderBy        string   `json:"orderBy" query:"orderBy"`
	Search         string   `json:"search" query:"search"`
}

type Meta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type GetProductsModel struct {
	Products []ProductModel `json:"products"`
	Meta     Meta           `json:"meta"`
}