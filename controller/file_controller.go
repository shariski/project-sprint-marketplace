package controller

import (
	"path"
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/middleware"
	"project-sprint-marketplace/service"
	"slices"

	"github.com/gofiber/fiber/v2"
)

type FileController struct {
	service.FileService
	configuration.Config
}

func NewFileController(fileService *service.FileService, config configuration.Config) *FileController {
	return &FileController{FileService: *fileService, Config: config}
}

func (controller FileController) Route(app *fiber.App) {
	app.Post("/v1/image", middleware.ValidateJWT(controller.Config), controller.Upload)
}

func (controller FileController) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	exception.PanicLogging(err)

	if file.Size > 2*1024*1024 || file.Size < 10*1024 {
		exception.PanicLogging(exception.ValidationError{
			Message: "File size must no more than 2MB, no less than 10KB",
		})
	}

	if ext := path.Ext(file.Filename); !slices.Contains([]string{".jpg", ".jpeg"}, ext) {
		exception.PanicLogging(exception.ValidationError{
			Message: "File type must be .jpg or .jpeg",
		})
	}

	result := controller.FileService.Upload(c.Context(), file)

	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"imageUrl": result,
	})
}