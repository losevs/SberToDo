package server

import (
	"sber/routes"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()
	routes.SetupRoutes(app)
	return app
}
