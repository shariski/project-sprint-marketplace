package service

import (
	"context"
	"project-sprint-marketplace/model"
)

type UserService interface {
	Authentication(ctx context.Context, model model.UserModel) model.UserGetModel
	Create(ctx context.Context, model model.UserModel) model.UserGetModel
}