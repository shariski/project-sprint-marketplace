package impl

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/repository"
)

type bankAccountRepositoryImpl struct{}

func NewBankAccountRepositoryImpl() repository.BankAccountRepository {
	return &bankAccountRepositoryImpl{}
}

func (bankAccountRepository *bankAccountRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, bankAccount entity.BankAccount) entity.BankAccount {
	SQL := `
		INSERT INTO bank_accounts (user_id, bank_name, bank_account_name, bank_account_number, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		returning id;
	`
	timeNow := common.GetDateNowUTCFormat()
	err := tx.QueryRowContext(ctx, SQL, &bankAccount.UserId, &bankAccount.BankName, &bankAccount.BankAccountName, &bankAccount.BankAccountNumber, &timeNow, &timeNow).Scan(&bankAccount.Id)
	exception.PanicLogging(err)
	return bankAccount
}

func (bankAccountRepository *bankAccountRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, bankAccount entity.BankAccount) entity.BankAccount {
	SQL := `
		UPDATE bank_accounts
		SET bank_name = $1, bank_account_name = $2, bank_account_number = $3, updated_at = $4
		WHERE id = $5
		AND user_id = $6;
	`
	timeNow := common.GetDateNowUTCFormat()
	res, err := tx.ExecContext(ctx, SQL,
		&bankAccount.BankName, &bankAccount.BankAccountName, &bankAccount.BankAccountNumber, &timeNow, &bankAccount.Id, &bankAccount.UserId)
	exception.PanicLogging(err)
	rowsAffected, err := res.RowsAffected()
	exception.PanicLogging(err)
	if rowsAffected < 1 {
		exception.PanicLogging(exception.NotFoundError{
			Message: "Bank account not found",
		})
	}
	return bankAccount
}

func (bankAccountRepository *bankAccountRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, id int) []entity.BankAccount {
	SQL := `
		SELECT id, bank_name, bank_account_name, bank_account_number
		FROM bank_accounts
		WHERE user_id = $1;
	`
	rows, err := tx.QueryContext(ctx, SQL, &id)
	exception.PanicLogging(err)
	defer rows.Close()

	var banks []entity.BankAccount

	for rows.Next() {
		var bank entity.BankAccount
		err := rows.Scan(&bank.Id, &bank.BankName, &bank.BankAccountName, &bank.BankAccountNumber)
		exception.PanicLogging(err)
		banks = append(banks, bank)
	}
	err = rows.Err()
	exception.PanicLogging(err)

	return banks
}

func (bankAccountRepository *bankAccountRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int, userId int) {
	SQL := `
		DELETE from bank_accounts
		WHERE id = $1
		AND user_id = $2;
	`
	res, err := tx.ExecContext(ctx, SQL, &id, &userId)
	exception.PanicLogging(err)

	rowsAffected, err := res.RowsAffected()
	exception.PanicLogging(err)

	if rowsAffected < 1 {
		exception.PanicLogging(exception.NotFoundError{
			Message: "Bank account not found",
		})
	}
}

func (bankAccountRepository *bankAccountRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int, userId int) entity.BankAccount {
	SQL := `
		SELECT id
		FROM bank_accounts
		WHERE id = $1
		AND user_id = $2;
	`
	var bank entity.BankAccount
	err := tx.QueryRowContext(ctx, SQL, &id).Scan(&bank.Id, &userId)
	if err == sql.ErrNoRows {
		exception.PanicLogging(exception.BadRequestError{
			Message: "Bank account not found",
		})
	}
	exception.PanicLogging(err)
	return bank
}
