package main

import (
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/controller"
	"project-sprint-marketplace/exception"
	repository "project-sprint-marketplace/repository/impl"
	service "project-sprint-marketplace/service/impl"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := configuration.New()
	database := configuration.NewDatabase(config)

	userRepository := repository.NewUserRepositoryImpl(database)

	userService := service.NewUserServiceImpl(&userRepository)

	userController := controller.NewUserController(&userService, config)

	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())

	userController.Route(app)

	err := app.Listen(":8000")
	exception.PanicLogging(err)
}
