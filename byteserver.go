package main

import (
	"github.com/gofiber/fiber/v2"
	"byteserver/pkg/projects"
)

func main() {
	port := ":3000"
	app := fiber.New()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BYTE Server is running!")
	})
	app.Mount("/projects", projects.Projects())

	app.Listen(port)
}
