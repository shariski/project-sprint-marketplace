package repository

import (
	"context"
	"project-sprint-marketplace/entity"
)

type ProductRepository interface {
	Insert(ctx context.Context, product entity.Product) (int, error)
}
