package main

import (
	"project-sprint-marketplace/configuration"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configuration.New()
	_ = configuration.NewDatabase(config)

	app := fiber.New()
	app.Listen(":8000")
}
