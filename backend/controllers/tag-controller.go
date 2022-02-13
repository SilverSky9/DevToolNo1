package controllers

import (
	"daeng-market/services"

	"github.com/gofiber/fiber/v2"
)

func TagRoute(r fiber.Router) {

	r.Get("/getall", func(c *fiber.Ctx) error {
		resp, err := services.GetAllTag()
		if err != nil {
			return c.Status(500).Send([]byte(err.Error()))
		}

		return c.Status(200).JSON(resp)
	})

}
