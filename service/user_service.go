package service

import (
	"context"
	"project-sprint-marketplace/model"
)

type UserService interface {
	Create(ctx context.Context, model model.UserModel) model.UserGetModel
}
