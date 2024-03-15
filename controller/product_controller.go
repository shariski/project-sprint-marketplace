package controller

import (
	"encoding/json"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/middleware"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service.ProductService
	service.PaymentService
	configuration.Config
}

func NewProductController(productService *service.ProductService, paymentService *service.PaymentService, config configuration.Config) *ProductController {
	return &ProductController{ProductService: *productService, PaymentService: *paymentService, Config: config}
}

func (controller ProductController) Route(app *fiber.App) {
	app.Post("/v1/product", middleware.ValidateJWT(controller.Config), controller.Create)
	app.Patch("v1/product/:id", controller.Update)
	app.Delete("/v1/product/:id", controller.DeleteById)
	app.Get("/v1/product/:id", controller.GetById)
	app.Post("v1/product/:id/stock", controller.UpdateStock)
	app.Post("v1/product/:id/buy", middleware.ValidateJWT(controller.Config), controller.CreatePayment)
}

func (controller ProductController) Create(c *fiber.Ctx) error {
	var request model.ProductCreateModel
	err := c.BodyParser(&request)
	request.UserId = c.Locals("userId").(int)

	exception.PanicLogging(err)

	errors := common.ValidateInput(request)

	if errors != nil {
		panic(exception.ValidationError{
			Message: errors.Error(),
		})
	}

	_ = controller.ProductService.Create(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "product added successfully",
	})
}

func (controller ProductController) Update(c *fiber.Ctx) error {
	var request model.ProductUpdateModel
	err := c.BodyParser(&request)

	exception.PanicLogging(err)

	productId := c.Params("id")

	request.Id, err = strconv.Atoi(productId)

	exception.PanicLogging(err)

	errors := common.ValidateInput(request)

	if errors != nil {
		panic(exception.ValidationError{
			Message: errors.Error(),
		})
	}

	_ = controller.ProductService.Update(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "product updated successfully",
	})
}

func (controller ProductController) DeleteById(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))

	exception.PanicLogging(err)

	controller.ProductService.DeleteById(c.Context(), productId)

	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "product deleted successfully",
	})
}

func (controller ProductController) GetById(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))

	exception.PanicLogging(err)

	result := controller.ProductService.FindById(c.Context(), productId)

	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "ok",
		Data:    result,
	})
}

func (controller ProductController) UpdateStock(c *fiber.Ctx) error {
	var request model.UpdateStockModel
	body := c.Body()

	err := json.Unmarshal(body, &request)

	exception.PanicLogging(err)

	productId := c.Params("id")

	request.Id, err = strconv.Atoi(productId)

	exception.PanicLogging(err)

	errors := common.ValidateInput(request)

	if errors != nil {
		panic(exception.ValidationError{
			Message: errors.Error(),
		})
	}

	_ = controller.ProductService.UpdateStockById(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "stock updated successfully",
	})
}

func (controller ProductController) CreatePayment(c *fiber.Ctx) error {
	var request model.PaymentModel
	body := c.Body()
	err := json.Unmarshal(body, &request)
	exception.PanicLogging(err)
	productId := c.Params("id")
	request.ProductId = productId

	userId := c.Locals("userId").(int)

	errors := common.ValidateInput(request)
	if errors != nil {
		panic(exception.ValidationError{
			Message: errors.Error(),
		})
	}

	_ = controller.PaymentService.Create(c.Context(), request, userId)

	return c.Status(fiber.StatusOK).JSON(model.ResponseFormat{
		Message: "payment processed successfully",
	})
}
