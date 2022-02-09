package controllers

import (
	"daeng-market/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		msg, _ := services.GetMessage()
		return c.Status(200).JSON(msg)
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		msg, _ := services.GetHealthCheck()
		return c.Status(200).JSON(msg)
	})
	app.Get("/post", func(c *fiber.Ctx) error {
		msg, _ := services.GetAllPost()
		return c.Status(200).JSON(msg)
	})
	// app.Get("/post/:id", func(c *fiber.Ctx) error {
	// 	msg := services.GetPostById(c.Params("id"))
	// 	return c.Status(200).JSON(msg)
	// })
	app.Get("/user", func(c *fiber.Ctx) error {
		msg, _ := services.GetUser()
		fmt.Println(msg, "From controller")
		return c.Status(200).JSON(msg)
	})

}
