package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/repository"
	"strings"

	"github.com/lib/pq"
)

type userRepositoryImpl struct {
	*sql.DB
}

func NewUserRepositoryImpl(DB *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

func (userRepository *userRepositoryImpl) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	sql := `
		INSERT INTO users (username, name, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING username, name;
	`
	var insertedUser entity.User
	if err := userRepository.DB.QueryRow(sql,
		&user.Username, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt).Scan(&insertedUser.Username, &insertedUser.Name); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return entity.User{}, errors.New("duplicate key")
		}
	}
	return insertedUser, nil
}

func (userRepository *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	query := `
		SELECT username, name, password
		FROM users
		WHERE username = $1
	`
	var user entity.User
	if err := userRepository.DB.QueryRow(query, &username).Scan(&user.Username, &user.Name, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, errors.New("no rows")
		}
	}
	return user, nil
}

func (userRepository *userRepositoryImpl) FindByProductId(ctx context.Context, productId int) model.SellerModel {
	query := `
		SELECT u.name AS seller_name, COALESCE(SUM(py.quantity), 0) AS total_products_sold, ARRAY_AGG(
			JSON_BUILD_OBJECT(
				'bank_account_id', COALESCE(ba.id::VARCHAR, ''),
				'bank_name', COALESCE(ba.bank_name, ''),
				'bank_account_name', COALESCE(ba.bank_account_name, ''),
				'bank_account_number', COALESCE(ba.bank_account_number, '')
			)
		) AS bank_accounts
		FROM users u
		JOIN products p ON u.id = p.user_id
		LEFT JOIN payments py ON p.id = py.product_id
		LEFT JOIN bank_accounts ba ON u.id = ba.user_id
		WHERE p.id = $1
		GROUP BY u.id, u.name;
	`

	var seller model.SellerModel
	var bankAccountsString pq.StringArray

	if err := userRepository.DB.QueryRow(query, &productId).Scan(&seller.Name, &seller.ProductSoldTotal, &bankAccountsString); err != nil {
		if err == sql.ErrNoRows {
			panic(exception.NotFoundError{
				Message: "seller not found",
			})
		} else {
			exception.PanicLogging(err)
		}
	}

	datas := []string(bankAccountsString)

	var bankAccounts []model.BankAccount

	for _, data := range datas {
		var bankAccount model.BankAccount
		
		err := json.Unmarshal([]byte(data), &bankAccount)
		exception.PanicLogging(err)

		bankAccounts = append(bankAccounts, bankAccount)
	}

	seller.BankAccounts = bankAccounts

	return seller
}