package controllers

import (
	model "daeng-market/models"
	"daeng-market/services"
	"github.com/gofiber/fiber/v2"
)

func PostRoute(r fiber.Router) {

	r.Post("/create", func(c *fiber.Ctx) error {
		var post model.Post
		err := c.BodyParser(&post)
		if err != nil {
			return c.SendStatus(400)
		}

		var tag model.Tag
		err = c.BodyParser(&tag)
		if err != nil {
			return c.SendStatus(400)
		}

		var pin model.Pin
		err = c.BodyParser(&pin)
		if err != nil {
			return c.SendStatus(400)
		}

		var user model.User
		err = c.BodyParser(&user)
		if err != nil {
			return c.SendStatus(400)
		}

		err = services.CreatePost(post, tag, user, pin)
		if err != nil {
			return c.Status(500).JSON(err)
		}

		return c.Status(201).JSON("Create post success!")
	})
}
