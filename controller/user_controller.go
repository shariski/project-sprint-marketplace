package controller

import (
	"project-sprint-marketplace/configuration"
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
	app.Post("/v1/user/register", controller.Create)
}

func (controller UserController) Create(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	if err != nil {
		panic(err)
	}
	result := controller.UserService.Create(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.ResponseFormat{
		Message: "User registered succesfully",
		Data:    result,
	})
}
