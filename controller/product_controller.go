package controller

import (
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service.ProductService
	configuration.Config
}

func NewProductController(productService *service.ProductService, config configuration.Config) *ProductController {
	return &ProductController{ProductService: *productService, Config: config}
}

func (controller ProductController) Route(app *fiber.App) {
	app.Post("/v1/product", controller.Create)
	app.Patch("v1/product/:id", controller.Update)
}

func (controller ProductController) Create(c *fiber.Ctx) error {
	var request model.ProductCreateModel
	err := c.BodyParser(&request)
	request.UserId = 1 //hardcoded
	
	if err != nil {
		panic(err)
	}

	errors := common.ValidateInput(request)

	if errors != nil {
		panic(exception.ValidationError{
			Message: errors.Error(),
		})
	}
	
	_ = controller.ProductService.Create(c.Context(), request)

	return c.Status(fiber.StatusCreated).JSON(model.ResponseFormat{
		Message: "product added successfully",
	})
}

func (controller ProductController) Update(c *fiber.Ctx) error {
	var request model.ProductUpdateModel
	err := c.BodyParser(&request)
	
	if err != nil {
		panic(err)
	}
	
	productId := c.Params("id")

	request.Id, err = strconv.Atoi(productId)
	
	if err != nil {
		panic(err)
	}

	errors := common.ValidateInput(request)

	if errors != nil {
		panic(exception.ValidationError{
			Message: errors.Error(),
		})
	}
	
	_ = controller.ProductService.Update(c.Context(), request)

	return c.Status(fiber.StatusCreated).JSON(model.ResponseFormat{
		Message: "product updated successfully",
	})
}
