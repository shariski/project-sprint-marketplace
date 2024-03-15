package model

type PaymentModel struct {
	ProductId            string `json:"productId"`
	BankAccountId        string `json:"bankAccountId"`
	PaymentProofImageUrl string `json:"paymentProofImageUrl"`
	Quantity             int    `json:"quantity"`
}
