package entity

type Payment struct {
	Id                   int    `sql:"id"`
	ProductId            int    `sql:"product_id"`
	BankAccountId        int    `sql:"bank_account_id"`
	PaymentProofImageUrl string `sql:"payment_proof_image_url"`
	Quantity             int    `sql:"quantity"`
	CreatedAt            string `sql:"created_at"`
	UpdatedAt            string `sql:"updated_at"`
}
