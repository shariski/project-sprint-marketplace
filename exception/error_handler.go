package exception

import (
	"project-sprint-marketplace/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		// var messages []map[string]interface{}

		// errJson := json.Unmarshal([]byte(data), &messages)
		// PanicLogging(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseErrorFormat{
			Message: "Bad Request",
			Errors: data,
		})
	}

	_, conflictError := err.(ConflictError)
	if conflictError {
		return ctx.Status(fiber.StatusConflict).JSON(model.ResponseErrorFormat{
			Message: "Conflict",
		})
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(model.ResponseErrorFormat{
			Message: "Not Found",
		})
	}

	_, forbiddenError := err.(ForbiddenError)
	if forbiddenError {
		return ctx.Status(fiber.StatusForbidden).JSON(model.ResponseErrorFormat{
			Message: "Forbidden",
		})
	}

	_, badRequestError := err.(BadRequestError)
	if badRequestError {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseErrorFormat{
			Message: "Bad Request",
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseErrorFormat{
		Message: "General Error",
	})
}
