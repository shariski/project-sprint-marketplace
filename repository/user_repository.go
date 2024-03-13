package repository

import (
	"context"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/model"
)

type UserRepository interface {
	// Authentication(ctx context.Context, username string, password string) (entity.User, error)
	Insert(ctx context.Context, user entity.User) (entity.User, error)
	FindByUsername(ctx context.Context, username string) (entity.User, error)
	FindByProductId(ctx context.Context, productId int) model.SellerModel
}
