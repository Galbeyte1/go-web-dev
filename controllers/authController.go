package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
	// localhost:port you should see what is sent
}

func Register(c *fiber.Ctx) error {
	// requests can be intercepted to this function and sent back
	// path specific /register
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	return c.JSON(data)
}