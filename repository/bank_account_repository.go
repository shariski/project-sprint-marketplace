package repository

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/entity"
)

type BankAccountRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, bankAccount entity.BankAccount) entity.BankAccount
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []entity.BankAccount
	Update(ctx context.Context, tx *sql.Tx, bankAccount entity.BankAccount) entity.BankAccount
	Delete(ctx context.Context, tx *sql.Tx, id int, userId int)
	FindById(ctx context.Context, tx *sql.Tx, id int, userId int) entity.BankAccount
}
