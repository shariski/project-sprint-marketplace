package main

import (
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/controller"
	"project-sprint-marketplace/exception"
	repository "project-sprint-marketplace/repository/impl"
	service "project-sprint-marketplace/service/impl"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := configuration.New()
	database := configuration.NewDatabase(config)
	storage := configuration.NewStorage(config)

	userRepository := repository.NewUserRepositoryImpl(database)
	productRepository := repository.NewProductRepositoryImpl()
	tagRepository := repository.NewTagRepositoryImpl()
	paymentRepository := repository.NewPaymentRepositoryImpl()
	bankAccountRepository := repository.NewBankAccountRepositoryImpl()

	userService := service.NewUserServiceImpl(&userRepository)
	productService := service.NewProductServiceImpl(database, &productRepository, &tagRepository, &userRepository)
	paymentService := service.NewPaymentServiceImpl(database, &paymentRepository, &productRepository, &bankAccountRepository)
	bankAccountService := service.NewBankAccountRepositoryImpl(database, &bankAccountRepository)
	fileService := service.NewFileServiceImpl(storage)

	userController := controller.NewUserController(&userService, config)
	productController := controller.NewProductController(&productService, &paymentService, config)
	bankAccountController := controller.NewBankAccountController(&bankAccountService, config)
	fileController := controller.NewFileController(&fileService, config)

	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(logger.New())

	userController.Route(app)
	productController.Route(app)
	bankAccountController.Route(app)
	fileController.Route(app)

	err := app.Listen(":8000")
	exception.PanicLogging(err)
}
