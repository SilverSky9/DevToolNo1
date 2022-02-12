package main

import (
	"daeng-market/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	controllers.Route(app)
	app.Listen(":3000")

}
