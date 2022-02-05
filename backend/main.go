package main

import (
	"daeng-market/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON("test")
	// })

	controllers.Route(app)

	app.Listen(":3000")
}
