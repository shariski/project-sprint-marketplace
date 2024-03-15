package entity

type BankAccount struct {
	Id                int    `sql:"id"`
	UserId            int    `sql:"user_id"`
	BankName          string `sql:"bank_name"`
	BankAccountName   string `sql:"bank_account_name"`
	BankAccountNumber string `sql:"bank_account_number"`
	CreatedAt         string `sql:"created_at"`
	UpdatedAt         string `sql:"updated_at"`
}
