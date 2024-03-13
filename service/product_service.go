package service

import (
	"context"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/model"
)

type ProductService interface {
	Create(ctx context.Context, model model.ProductCreateModel) entity.Product
	Update(ctx context.Context, model model.ProductUpdateModel) entity.Product
}
