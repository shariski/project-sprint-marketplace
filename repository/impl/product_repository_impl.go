package impl

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/repository"

	"github.com/lib/pq"
)

type productRepositoryImpl struct {}

func NewProductRepositoryImpl() repository.ProductRepository {
	return &productRepositoryImpl{}
}

func (productRepository *productRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id int) model.ProductModel {
	SQL := `
		SELECT p.id::varchar AS product_id, p.name, p.price, p.image_url, p.stock, p.condition, ARRAY_AGG(t.name) AS tags, p.is_purchaseable, COALESCE(SUM(py.quantity), 0) AS purchase_count
		FROM products p
		LEFT JOIN tags t ON p.id = t.product_id
		LEFT JOIN payments py ON p.id = py.product_id
		WHERE p.id = $1
		GROUP BY p.id, p.name, p.price, p.image_url, p.stock, p.condition, p.is_purchaseable;
	`

	var product model.ProductModel
	var tags pq.StringArray
		
	err := db.QueryRow(SQL, id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.ImageUrl,
		&product.Stock,
		&product.Condition,
		&tags,
		&product.IsPurchaseable,
		&product.PurchaseCount,
	)

	product.Tags = []string(tags)

	if err != nil {
		if err == sql.ErrNoRows {
			panic(exception.NotFoundError{
				Message: "product not found",
			})
		} else {
			exception.PanicLogging(err)
		}
	}

	return product
}

func (productRepository *productRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product {
	SQL := `
		INSERT INTO products (user_id, name, price, image_url, stock, condition, is_purchaseable, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id;
	`

	timeNow := common.GetDateNowUTCFormat()

	err := tx.QueryRowContext(ctx, SQL,
		&product.UserId, &product.Name, &product.Price, &product.ImageUrl, &product.Stock, &product.Condition, &product.IsPurchaseable, timeNow, timeNow).Scan(&product.Id)

	exception.PanicLogging(err)

	product.CreatedAt = timeNow
	product.UpdatedAt = timeNow

	return product
}

func (productRepository *productRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product {
	SQL := `
		UPDATE products SET name = $1, price = $2, image_url = $3, condition = $4, is_purchaseable = $5, updated_at = $6
		WHERE id = $7;
		`

	res, err := tx.ExecContext(ctx, SQL,
		&product.Name, &product.Price, &product.ImageUrl, &product.Condition, &product.IsPurchaseable, common.GetDateNowUTCFormat(), &product.Id)
		
	exception.PanicLogging(err)

	rowsAffected, err := res.RowsAffected()

	exception.PanicLogging(err)

	if rowsAffected < 1 {
		exception.PanicLogging(exception.NotFoundError{
			Message: "Product not found",
		})
	}

	return product
}

func (productRepository *productRepositoryImpl) DeleteByProductId(ctx context.Context, tx *sql.Tx, id int) {
	SQL := `
		DELETE FROM products
		WHERE id = $1;
		`

		res, err := tx.ExecContext(ctx, SQL, id)

		exception.PanicLogging(err)

		rowsAffected, err := res.RowsAffected()

		exception.PanicLogging(err)

		if rowsAffected < 1 {
			exception.PanicLogging(exception.NotFoundError{
				Message: "Product not found",
			})
		}
}

func (productRepository *productRepositoryImpl) UpdateStock(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product {
	SQL := `
		UPDATE products SET stock = $1
		WHERE id = $2;
		`

	res, err := tx.ExecContext(ctx, SQL,
		&product.Stock, &product.Id)
	exception.PanicLogging(err)

	rowsAffected, err := res.RowsAffected()
	exception.PanicLogging(err)

	if rowsAffected < 1 {
		exception.PanicLogging(exception.NotFoundError{
			Message: "Product not found",
		})
	}

	return product
}