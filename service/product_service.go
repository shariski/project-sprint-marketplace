package service

import (
	"context"
	"project-sprint-marketplace/model"
)

type ProductService interface {
	Create(ctx context.Context, model model.ProductCreateModel) error
	Update(ctx context.Context, model model.ProductUpdateModel) error
}
