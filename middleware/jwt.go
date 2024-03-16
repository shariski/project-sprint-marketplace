package middleware

import (
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/model"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateJWT(config configuration.Config) func(*fiber.Ctx) error {
	secret := config.Get("JWT_SECRET")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			id := claims["id"].(float64)
			c.Locals("userId", int(id))
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if c.Get("Authorization") == "" {
				return c.Status(fiber.StatusUnauthorized).JSON(model.ResponseErrorFormat{
					Message: "Unauthorized",
				})
			}
			return c.Status(fiber.StatusForbidden).JSON(model.ResponseErrorFormat{
				Message: "Forbidden",
			})
		},
	})
}

func ValidateOptionalJWT(config configuration.Config) func(*fiber.Ctx) error {
	secret := config.Get("JWT_SECRET")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			id := claims["id"].(float64)
			c.Locals("userId", int(id))
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Locals("userId", nil)
			return c.Next()
		},
	})
}
