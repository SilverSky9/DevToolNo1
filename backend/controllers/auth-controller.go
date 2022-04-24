package controllers

import (
	services "daeng-market/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router) {

	r.Post("/register", func(c *fiber.Ctx) error {
		type pin_request struct {
			PostId int    `json:"post_id"`
			PinId  int    `json:"pin_id"`
			Pin    string `json:"pin"`
		}

		var pin pin_request
		err := c.BodyParser(&pin)
		if err != nil {
			return c.SendStatus(400)
		}

		status, err := services.RegisterNewToken(pin.Pin, pin.PostId, c)
		if status != 200 {
			return c.SendStatus(status)
		}

		return c.SendStatus(200)
	})

	r.Post("/verify/:post_id", func(c *fiber.Ctx) error {
		post_id := c.Params("post_id")

		token_verify := services.VerifyTokenCookies(c, post_id)
		if token_verify {
			resp, err := services.GetUserByPostId(post_id)
			if err != nil {
				return c.SendStatus(403)
			}

			return c.Status(200).JSON(resp)
		}

		return c.SendStatus(403)
	})

	r.Get("/extract/:token", func(c *fiber.Ctx) error {
		token := c.Params("token")

		resp, err := services.ExtractToken(token)
		if err != nil {
			return c.SendStatus(500)
		}

		fmt.Println(resp)

		return c.SendStatus(200)
	})

}
