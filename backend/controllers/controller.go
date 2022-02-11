package controllers

import (
	services "daeng-market/services"
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
	// app.Get("/post", func(c *fiber.Ctx) error {
	// 	msg, _ := services.GetAllPost()
	// 	return c.Status(200).JSON(msg)
	// })

	//make post entry point
	post_entry_point := app.Group("/post")
	PostRoute(post_entry_point)
	//make user entry point
	user_entry_point := app.Group("/user")
	UserRoute(user_entry_point)

	app.Get("/post/:id", func(c *fiber.Ctx) error {
		msg := services.GetPostById(c.Params("id"))
		return c.Status(200).JSON(msg)
	})
	app.Get("/post/tag/:tag1/:tag2", func(c *fiber.Ctx) error {
		msg, _ := postService.GetPostByTag(c.Params("tag1"), c.Params("tag2"))
		return c.Status(200).JSON(msg)
	})
	app.Get("/user", func(c *fiber.Ctx) error {
		msg, _ := postService.GetAllUser()
		fmt.Println(msg, "From controller")
		return c.Status(200).JSON(msg)
	})

}
