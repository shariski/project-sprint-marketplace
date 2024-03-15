package impl

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/repository"
	"project-sprint-marketplace/service"
	"strconv"
)

type paymentServiceImpl struct {
	*sql.DB
	repository.PaymentRepository
	repository.ProductRepository
	repository.BankAccountRepository
}

func NewPaymentServiceImpl(
	DB *sql.DB,
	paymentRepository *repository.PaymentRepository,
	productRepository *repository.ProductRepository,
	bankAccountRepository *repository.BankAccountRepository,
) service.PaymentService {
	return &paymentServiceImpl{
		DB:                    DB,
		PaymentRepository:     *paymentRepository,
		ProductRepository:     *productRepository,
		BankAccountRepository: *bankAccountRepository,
	}
}

func (paymentService *paymentServiceImpl) Create(ctx context.Context, data model.PaymentModel, userId int) entity.Payment {
	productId, err := strconv.Atoi(data.ProductId)
	exception.PanicLogging(err)
	bankAccountId, err := strconv.Atoi(data.BankAccountId)
	exception.PanicLogging(err)
	payment := entity.Payment{
		ProductId:            productId,
		BankAccountId:        bankAccountId,
		PaymentProofImageUrl: data.PaymentProofImageUrl,
		Quantity:             data.Quantity,
	}
	product := entity.Product{
		Id: productId,
	}

	tx, err := paymentService.DB.Begin()
	exception.PanicLogging(err)
	defer common.CommitOrRollback(tx)

	_ = paymentService.BankAccountRepository.FindById(ctx, tx, bankAccountId, userId)
	insertedPayment := paymentService.PaymentRepository.Insert(ctx, tx, payment)
	paymentService.ProductRepository.DecrementStock(ctx, tx, product, data.Quantity)

	return insertedPayment
}
