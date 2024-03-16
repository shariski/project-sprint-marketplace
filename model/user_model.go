package model

type UserModel struct {
	Username string `json:"username" validate:"required,min=5,max=15"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type UserLoginModel struct {
	Username string `json:"username" validate:"required,min=5,max=15"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type UserGetModel struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type AuthenticationModel struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type SellerModel struct {
	Name             string        `json:"name"`
	ProductSoldTotal int           `json:"productSoldTotal"`
	BankAccounts     []BankAccount `json:"bankAccounts"`
}
