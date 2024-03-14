package middleware

import (
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func ValidateJWT(config configuration.Config) func(*fiber.Ctx) error {
	secret := config.Get("JWT_SECRET")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		SuccessHandler: func(c *fiber.Ctx) error {
			authHeader := c.Get("Authorization")
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.Status(fiber.StatusUnauthorized).JSON(model.ResponseErrorFormat{
					Message: "Unauthorized",
				})
			}
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(model.ResponseErrorFormat{
				Message: "Unauthorized",
			})
		},
	})
}
