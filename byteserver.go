package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := ":3000"

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(port)
}
