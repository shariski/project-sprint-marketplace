package model

type BankAccount struct {
	UserId            int    `json:"-"`
	BankName          string `json:"bankName" validate:"required,min=5,max=15"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=5,max=15"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=5,max=15"`
}

type BankAccountGetModel struct {
	BankAccountId     string `json:"bankAccountId"`
	BankName          string `json:"bankName"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountNumber string `json:"bankAccountNumber"`
}
type BankAccountUpdateModel struct {
	Id                int    `json:"-"`
	UserId            int    `json:"-"`
	BankName          string `json:"bankName" validate:"required,min=5,max=15"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=5,max=15"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=5,max=15"`
}
