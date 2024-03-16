package controller

import (
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/middleware"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BankAccountController struct {
	service.BankAccountService
	configuration.Config
}

func NewBankAccountController(bankAccountService *service.BankAccountService, config configuration.Config) *BankAccountController {
	return &BankAccountController{BankAccountService: *bankAccountService, Config: config}
}

func (controller BankAccountController) Route(app *fiber.App) {
	app.Post("/v1/bank/account", middleware.ValidateJWT(controller.Config), controller.Create)
	app.Get("/v1/bank/account", middleware.ValidateJWT(controller.Config), controller.GetByUserId)
	app.Patch("/v1/bank/account/:id", middleware.ValidateJWT(controller.Config), controller.Update)
	app.Delete("/v1/bank/account/:id", middleware.ValidateJWT(controller.Config), controller.Delete)
}

func (controller BankAccountController) Create(c *fiber.Ctx) error {
	var request model.BankAccount
	err := c.BodyParser(&request)
	exception.PanicLogging(err)
	request.UserId = c.Locals("userId").(int)

	errors := common.ValidateInput(request)
	if errors != nil {
		panic(exception.ValidationError{
			Message: errors.Error(),
		})
	}
	_ = controller.BankAccountService.Create(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "account added successfully",
	})
}

func (controller BankAccountController) GetByUserId(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)

	banks := controller.BankAccountService.FindByUserId(c.Context(), userId)
	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "success",
		Data:    banks,
	})
}

func (controller BankAccountController) Update(c *fiber.Ctx) error {
	var request model.BankAccountUpdateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)
	id := c.Params("id")
	bankId, err := strconv.Atoi(id)
	exception.PanicLogging(err)
	userId := c.Locals("userId").(int)
	request.Id = bankId
	request.UserId = userId

	errors := common.ValidateInput(request)
	if errors != nil {
		panic(exception.ValidationError{
			Message: errors.Error(),
		})
	}

	controller.BankAccountService.Update(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "account updated successfuly",
	})
}

func (controller BankAccountController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	bankId, err := strconv.Atoi(id)
	exception.PanicLogging(err)

	userId := c.Locals("userId").(int)

	controller.BankAccountService.Delete(c.Context(), bankId, userId)
	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "account deleted successfuly",
	})
}
