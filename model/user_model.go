package model

type UserModel struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
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
