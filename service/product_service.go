package service

import (
	"context"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/model"
)

type ProductService interface {
	Create(ctx context.Context, model model.ProductCreateModel) entity.Product
	Update(ctx context.Context, model model.ProductUpdateModel) entity.Product
	DeleteById(ctx context.Context, id int, userId int)
	FindById(ctx context.Context, id int) model.GetProductModel
	UpdateStock(ctx context.Context, data model.UpdateStockModel) entity.Product
}
