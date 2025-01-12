package main

import (
	"byteserver/pkg/database"
	"byteserver/pkg/projects"

	"github.com/gofiber/fiber/v2"
)

func main() {	
	port := ":3000"
	app := fiber.New()
	database.InitDB()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BYTE Server is running!")
	})
	app.Mount("/projects", projects.Projects())

	app.Listen(port)
}
