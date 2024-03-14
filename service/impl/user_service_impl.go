package impl

import (
	"context"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/repository"
	"project-sprint-marketplace/service"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repository.UserRepository
}

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

func (userService *userServiceImpl) Authentication(ctx context.Context, userModel model.UserModel) model.AuthenticationModel {
	userResult, err := userService.UserRepository.FindByUsername(ctx, userModel.Username)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(userModel.Password))
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Incorrect password",
		})
	}

	return model.AuthenticationModel{
		Id:       userResult.Id,
		Username: userResult.Username,
		Name:     userResult.Name,
	}
}

func (userService *userServiceImpl) Create(ctx context.Context, userModel model.UserModel) model.AuthenticationModel {
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
	return model.AuthenticationModel{
		Id:       userResult.Id,
		Username: userResult.Username,
		Name:     userResult.Name,
	}
}
