package main

import (
	"byteserver/pkg/database"
	"byteserver/pkg/projects"
	"byteserver/pkg/mongo"
	"github.com/gofiber/fiber/v2"
	"byteserver/pkg/applications"
)

func main() {	
	port := ":3000"
	app := fiber.New()
	database.InitDB()
	mongodb.Connect()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BYTE Server is running!")
	})
	app.Mount("/projects", projects.Projects())
	app.Mount("/applications", applications.ApplicationRouting())

	app.Listen(port)
}
