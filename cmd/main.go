package main

import (
	"sber/database"
	"sber/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	SetupRoutes(app)
	app.Listen(":80")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Show)
	app.Post("/add", handlers.Add)
}
