package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

var app *fiber.App = fiber.New()
var api *fiber.App = fiber.New()

func main() {
	app.Static("/", "./frontend/build")
	api.Post('login')

	app.Mount("/api", api)
	err := app.Listen(port)
	if err != nil {
		log.Fatal("Server exited with error message: ", err)
	}
}
