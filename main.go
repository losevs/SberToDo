package main

import (
	"sber/routes"
	"sber/server"
)

func main() {
	app := server.New()
	routes.SetupRoutes(app)
	app.Listen(":80")
}
