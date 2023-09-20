package main

import (
	"sber/database"
	_ "sber/docs"
	"sber/handlers"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	SetupRoutes(app)
	app.Listen(":80")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/show", handlers.Show)
	app.Get("/flag/:title", handlers.ChangeFlag)
	app.Get("/true", handlers.TruePag)
	app.Get("/false", handlers.FalsePag)
	app.Get("/date/:flag", handlers.FlagAsc)

	app.Post("/add", handlers.Add)

	app.Delete("/del/:title", handlers.Del)

	app.Patch("/change/:title", handlers.PatchToDo)
}
