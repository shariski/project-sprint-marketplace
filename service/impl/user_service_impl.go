package impl

import (
	"context"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/repository"
	"project-sprint-marketplace/service"
)

type userServiceImpl struct {
	repository.UserRepository
}

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

func (userService *userServiceImpl) Create(ctx context.Context, userModel model.UserModel) model.UserGetModel {
	user := entity.User{
		Username:  userModel.Username,
		Name:      userModel.Name,
		Password:  userModel.Password,
		CreatedAt: common.GetDateNowUTCFormat(),
		UpdatedAt: common.GetDateNowUTCFormat(),
	}
	userResult, err := userService.UserRepository.Insert(ctx, user)
	if err != nil {
		panic(exception.ConflictError{
			Message: "Username exists",
		})
	}
	return model.UserGetModel{
		Username: userResult.Username,
		Name:     userResult.Name,
	}
}
