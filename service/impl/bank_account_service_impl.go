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

type bankAccountServiceImpl struct {
	*sql.DB
	repository.BankAccountRepository
}

func NewBankAccountRepositoryImpl(
	DB *sql.DB,
	bankAccountRepository *repository.BankAccountRepository,
) service.BankAccountService {
	return &bankAccountServiceImpl{
		DB:                    DB,
		BankAccountRepository: *bankAccountRepository,
	}
}

func (bankAccountService *bankAccountServiceImpl) Create(ctx context.Context, bank model.BankAccount) model.BankAccountGetModel {
	bankAccount := entity.BankAccount{
		UserId:            bank.UserId,
		BankName:          bank.BankName,
		BankAccountName:   bank.BankAccountName,
		BankAccountNumber: bank.BankAccountNumber,
	}
	tx, err := bankAccountService.DB.Begin()
	exception.PanicLogging(err)
	defer common.CommitOrRollback(tx)

	insertedBank := bankAccountService.Insert(ctx, tx, bankAccount)

	return model.BankAccountGetModel{
		BankName:          insertedBank.BankName,
		BankAccountName:   insertedBank.BankAccountName,
		BankAccountNumber: insertedBank.BankAccountNumber,
	}
}

func (bankAccountService *bankAccountServiceImpl) FindByUserId(ctx context.Context, userId int) []model.BankAccountGetModel {
	tx, err := bankAccountService.DB.Begin()
	exception.PanicLogging(err)
	defer common.CommitOrRollback(tx)

	var banks []model.BankAccountGetModel
	res := bankAccountService.BankAccountRepository.FindByUserId(ctx, tx, userId)
	for _, bank := range res {
		banks = append(banks, model.BankAccountGetModel{
			BankAccountId:     strconv.Itoa(bank.Id),
			BankName:          bank.BankName,
			BankAccountName:   bank.BankAccountName,
			BankAccountNumber: bank.BankAccountNumber,
		})
	}

	return banks
}

func (bankAccountService *bankAccountServiceImpl) Update(ctx context.Context, data model.BankAccountUpdateModel) {
	tx, err := bankAccountService.DB.Begin()
	exception.PanicLogging(err)
	defer common.CommitOrRollback(tx)

	_ = bankAccountService.BankAccountRepository.Update(ctx, tx, entity.BankAccount{
		Id:                data.Id,
		UserId:            data.UserId,
		BankName:          data.BankName,
		BankAccountName:   data.BankAccountName,
		BankAccountNumber: data.BankAccountNumber,
	})
}

func (bankAccountService *bankAccountServiceImpl) Delete(ctx context.Context, id int, userId int) {
	tx, err := bankAccountService.DB.Begin()
	exception.PanicLogging(err)
	defer common.CommitOrRollback(tx)

	bankAccountService.BankAccountRepository.Delete(ctx, tx, id, userId)
}
