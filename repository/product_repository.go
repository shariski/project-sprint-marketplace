package repository

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/model"
)

type ProductRepository interface {
	FindByFilters(ctx context.Context, db *sql.DB, filters model.ProductFilters) ([]model.ProductModel, int)
	FindById(ctx context.Context, db *sql.DB, id int) entity.Product
	FindByIdAggregated(ctx context.Context, db *sql.DB, id int) model.ProductModel
	Insert(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product
	Update(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product
	DeleteByProductId(ctx context.Context, tx *sql.Tx, id int)
	UpdateStock(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product
}
