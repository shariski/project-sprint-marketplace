package impl

import (
	"context"
	"database/sql"
	"fmt"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/repository"
	"strconv"

	"github.com/lib/pq"
)

type productRepositoryImpl struct{}

func NewProductRepositoryImpl() repository.ProductRepository {
	return &productRepositoryImpl{}
}

func (productRepository *productRepositoryImpl) FindByFilters(ctx context.Context, db *sql.DB, filters model.ProductFilters) ([]model.ProductModel, int) {
	SQL := `
		SELECT p.id::varchar AS product_id, p.name, p.price, p.image_url, p.stock, p.condition, ARRAY_AGG(t.name) AS tags, p.is_purchaseable, COALESCE(SUM(py.quantity), 0) AS purchase_count
		FROM products p
		LEFT JOIN tags t ON p.id = t.product_id
		LEFT JOIN payments py ON p.id = py.product_id
	`

	COUNT := `
		SELECT COUNT(*) as total
		FROM
	`

	FILTERS := ""

	if filters.UserOnly {
		FILTERS += "p.user_id = " + strconv.Itoa(filters.UserId) + " "
	}

	if len(filters.Tags) > 0 {
		if (FILTERS != "") {
			FILTERS += "AND "
		}

		FILTERS += "("
		
		for i, tag := range filters.Tags {
			if i > 0 {
				FILTERS += "OR t.name = '" + tag + "' "
			} else {
				FILTERS += "t.name = '" + tag + "' "
			}
		}

		FILTERS += ") "
	}

	if filters.Condition != "" {
		if (FILTERS != "") {
			FILTERS += "AND "
		}

		FILTERS += "p.condition = '" + filters.Condition + "' "
	}

	if !filters.ShowEmptyStock {
		if (FILTERS != "") {
			FILTERS += "AND "
		}

		FILTERS += "p.stock > 0 "
	}

	if filters.MaxPrice > 0 {
		if (FILTERS != "") {
			FILTERS += "AND "
		}

		FILTERS += "p.price <= '" + fmt.Sprintf("%f", filters.MaxPrice) + "' " 
	}

	if filters.MinPrice > 0 {
		if (FILTERS != "") {
			FILTERS += "AND "
		}

		FILTERS += "p.price >= '" + fmt.Sprintf("%f", filters.MinPrice) + "' " 
	}

	if filters.Search != "" {
		if (FILTERS != "") {
			FILTERS += "AND "
		}

		FILTERS += "p.name LIKE '%" + filters.Search + "%' "
	}

	if FILTERS != "" {
		SQL += "WHERE " + FILTERS
	}

	SQL += "GROUP BY p.id, p.name, p.price, p.image_url, p.stock, p.condition, p.is_purchaseable "
	COUNT += "(" + SQL + ") AS subquery;"

	if filters.SortBy != "" {
		if filters.SortBy == "date" {
			filters.SortBy = "p.created_at"
		} else {
			filters.SortBy = "p.price"
		}
		SQL += "ORDER BY " + filters.SortBy + " "
		
		if filters.OrderBy != "" {
			SQL += filters.OrderBy + " "
		} else {
			SQL += "ASC" + " "
		}
	}


	if filters.Limit > 0 {
		SQL += "LIMIT " + strconv.Itoa(filters.Limit) + " "
	}

	if filters.Offset > 0 {
		SQL += "OFFSET " + strconv.Itoa(filters.Offset) + " "
	}

	SQL += ";"

	var total int

	err := db.QueryRow(COUNT).Scan(&total)
	exception.PanicLogging(err)

	rows, err := db.Query(SQL)
	if err != nil {
			exception.PanicLogging(err)
	}
	defer rows.Close()

	products := []model.ProductModel{}
	for rows.Next() {
			var product model.ProductModel
			var tags pq.StringArray

			if err := rows.Scan(
				&product.Id,
				&product.Name,
				&product.Price,
				&product.ImageUrl,
				&product.Stock,
				&product.Condition,
				&tags,
				&product.IsPurchaseable,
				&product.PurchaseCount,
			); err != nil {
				exception.PanicLogging(err)
			}

			product.Tags = []string(tags)

			products = append(products, product)
	}

	if err := rows.Err(); err != nil {
			exception.PanicLogging(err)
	}

	return products, total
}

func (productRepository *productRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id int) entity.Product {
	SQL := `
		SELECT * FROM products WHERE id = $1;
	`

	var product entity.Product

	err := db.QueryRow(SQL, id).Scan(
		&product.Id,
		&product.UserId,
		&product.Name,
		&product.Price,
		&product.ImageUrl,
		&product.Stock,
		&product.Condition,
		&product.IsPurchaseable,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

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

func (productRepository *productRepositoryImpl) FindByIdAggregated(ctx context.Context, db *sql.DB, id int) model.ProductModel {
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
		UPDATE products SET stock = $1, updated_at = $2
		WHERE id = $3;
		`

	res, err := tx.ExecContext(ctx, SQL,
		&product.Stock, common.GetDateNowUTCFormat(), &product.Id)
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

func (productRepository *productRepositoryImpl) DecrementStock(ctx context.Context, tx *sql.Tx, product entity.Product, quantity int) {
	SQL := `
		UPDATE products
		SET stocks = stocks - $1
		WHERE id = $2
		AND stocks > 0;
	`
	res, err := tx.ExecContext(ctx, SQL, &quantity, &product.Id)
	exception.PanicLogging(err)

	rowsAffected, err := res.RowsAffected()
	exception.PanicLogging(err)

	if rowsAffected < 1 {
		exception.PanicLogging(exception.BadRequestError{
			Message: "Product not found or insufficient stocks",
		})
	}
}
