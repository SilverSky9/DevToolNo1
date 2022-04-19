package controllers

import (
	model "daeng-market/models"
	"daeng-market/services"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(r fiber.Router) {

	r.Post("/create", func(c *fiber.Ctx) error {
		//get post data
		var post model.Post
		err := c.BodyParser(&post)
		if err != nil {
			return c.SendStatus(400)
		}

		//get tag name data
		var tag model.Tag
		err = c.BodyParser(&tag)
		if err != nil {
			return c.SendStatus(400)
		}

		//get input pin data
		var pin model.Pin
		err = c.BodyParser(&pin)
		if err != nil {
			return c.SendStatus(400)
		}

		//get owner data
		var user model.User
		err = c.BodyParser(&user)
		if err != nil {
			return c.SendStatus(400)
		}

		//create post
		err = services.CreatePost(post, tag, user, pin)
		if err != nil {
			return c.Status(500).JSON(err)
		}

		return c.Status(201).JSON("Create post success!")
	})

	r.Get("/getbyid/:id", func(c *fiber.Ctx) error {
		msg, err := services.GetPostById(c.Params("id"))
		if err != nil {
			return c.Status(500).JSON(err.Error())
		}
		return c.Status(200).JSON(msg)
	})

	r.Get("/getbytag/:tag1/:tag2", func(c *fiber.Ctx) error {
		tag_1 := c.Params("tag1")
		tag_2 := c.Params("tag2")

		msg, err := services.GetPostByTag(tag_1, tag_2)
		if err != nil {
			return c.Status(500).JSON(err.Error())
		}
		return c.Status(200).JSON(msg)
	})

	r.Get("/all", func(c *fiber.Ctx) error {
		msg, err := services.GetAllPost()

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}
		return c.Status(200).JSON(msg)
	})

	r.Get("/searchbyname/:name", func(c *fiber.Ctx) error {
		c.Accepts("text/html") // "text/html"
		c.Accepts("json", "text")
		post_name := c.Params("name")

		//decode encoded for thai language
		decode_name, err := url.QueryUnescape(post_name)
		if err != nil {
			return c.Status(500).JSON("decode params error")
		}

		msg, err := services.GetPostByName(decode_name)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}
		return c.Status(200).JSON(msg)
	})

	r.Get("/geybymultitag/:tags", func(c *fiber.Ctx) error {
		tag_list := c.Params("tags")
		resp, err := services.GetPostByMultiTag(tag_list)
		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(resp)
	})

	// r.Post("/getitruser", func(c *fiber.Ctx) error {
	// 	var pin model.Pin

	// 	// resp, err := services

	// 	return c.Status(200).JSON()
	// })

}
