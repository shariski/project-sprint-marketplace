package controller

import (
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/middleware"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service.UserService
	configuration.Config
}

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/user/login", middleware.ValidateJWT(controller.Config), controller.Authentication)
	app.Post("/v1/user/register", controller.Create)
}

func (controller UserController) Authentication(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.UserService.Authentication(c.Context(), request)
	accessToken := common.GenerateToken(result.Id, controller.Config)
	userAndToken := map[string]interface{}{
		"username":    result.Username,
		"name":        result.Name,
		"accessToken": accessToken,
	}
	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "User logged successfully",
		Data:    userAndToken,
	})
}

func (controller UserController) Create(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	bcryptPassword := common.GenerateBcryptPassword(request.Password, controller.Config)
	request.Password = bcryptPassword

	result := controller.UserService.Create(c.Context(), request)

	accessToken := common.GenerateToken(result.Id, controller.Config)
	userAndToken := map[string]interface{}{
		"username":    result.Username,
		"name":        result.Name,
		"accessToken": accessToken,
	}

	return c.Status(fiber.StatusCreated).JSON(model.ResponseFormat{
		Message: "User registered succesfully",
		Data:    userAndToken,
	})
}
