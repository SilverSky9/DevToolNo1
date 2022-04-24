package controllers

import (
	model "daeng-market/models"
	services "daeng-market/services"

	"github.com/gofiber/fiber/v2"
)

// Users collection
type PostList []model.Post

func Route(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		msg, _ := services.GetMessage()
		return c.Status(200).JSON(msg)
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		msg, _ := services.GetHealthCheck()
		return c.Status(200).JSON(msg)
	})

	//make post entry point
	post_entry_point := app.Group("/post")
	PostRoute(post_entry_point)

	//make user entry point
	user_entry_point := app.Group("/user")
	UserRoute(user_entry_point)

	//make tag entry point
	tag_entry_point := app.Group("/tag")
	TagRoute(tag_entry_point)

	//make auth entry point
	auth_entry_point := app.Group("/auth")
	AuthRoute(auth_entry_point)

}
