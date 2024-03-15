package service

import (
	"context"
	"project-sprint-marketplace/model"
)

type BankAccountService interface {
	Create(ctx context.Context, bank model.BankAccount) model.BankAccountGetModel
	FindByUserId(ctx context.Context, id int) []model.BankAccountGetModel
	Update(ctx context.Context, bank model.BankAccountUpdateModel)
	Delete(ctx context.Context, id int, userId int)
}
