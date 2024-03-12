package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/repository"
)

type productRepositoryImpl struct {
	*sql.DB
}

func NewProductRepositoryImpl(DB *sql.DB) repository.ProductRepository {
	return &productRepositoryImpl{DB: DB}
}

func (productRepository *productRepositoryImpl) Insert(ctx context.Context, product entity.Product) (int, error) {
	sql := `
		INSERT INTO products (user_id, name, price, image_url, stock, condition, is_purchaseable, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id;
	`

	var productId int

	err := productRepository.DB.QueryRow(sql,
		&product.UserId, &product.Name, &product.Price, &product.ImageUrl, &product.Stock, &product.Condition, &product.IsPurchaseable, common.GetDateNowUTCFormat(), common.GetDateNowUTCFormat()).Scan(&productId)

	return productId, err
}

func (productRepository *productRepositoryImpl) Update(ctx context.Context, product entity.Product) error {
	checkSql := `SELECT name FROM products WHERE id = $1`

	var name string
	
	err := productRepository.DB.QueryRow(checkSql, &product.Id).Scan(&name)
	if (err != nil) {
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			panic(exception.NotFoundError{
				Message: "product not found",
			})
		} else {
			panic(err)
		}
	}


	sql := `
		UPDATE products SET name = $1, price = $2, image_url = $3, condition = $4, is_purchaseable = $5, updated_at = $6
		WHERE id = $7;
		`

	err = productRepository.DB.QueryRow(sql,
		&product.Name, &product.Price, &product.ImageUrl, &product.Condition, &product.IsPurchaseable, common.GetDateNowUTCFormat(), &product.Id).Err()

	return err
}
