package impl

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/repository"
)

type paymentRepositoryImpl struct{}

func NewPaymentRepositoryImpl() repository.PaymentRepository {
	return &paymentRepositoryImpl{}
}

func (paymentRepository *paymentRepositoryImpl) Insert(c context.Context, tx *sql.Tx, payment entity.Payment) entity.Payment {
	query := `
		INSERT INTO payments (product_id, bank_account_id, payment_proof_image_url, quantity, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	timeNow := common.GetDateNowUTCFormat()
	err := tx.QueryRowContext(c, query,
		&payment.ProductId, &payment.BankAccountId, &payment.PaymentProofImageUrl, &payment.Quantity, &timeNow, &timeNow).Scan(&payment.Id)
	exception.PanicLogging(err)

	return payment
}
