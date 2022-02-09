package main

import (
	"daeng-market/controllers"
	"daeng-market/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	services.Connect_DB()
	controllers.Route(app)

	app.Listen(":3000")

}
