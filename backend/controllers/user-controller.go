package controllers

import (
	model "daeng-market/models"
	"daeng-market/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router) {
	r.Post("/create", func(c *fiber.Ctx) error {
		var user model.User
		err := c.BodyParser(&user)
		if err != nil {
			return c.SendStatus(400)
		}

		_, err = services.CreateUser(user)
		return c.Status(201).JSON("Create user success!")
	})
}
