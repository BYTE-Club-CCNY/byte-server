package main

import (
	"byteserver/pkg/database"
	"byteserver/pkg/projects"
	"byteserver/pkg/mongo"
	"github.com/gofiber/fiber/v2"
	"byteserver/pkg/apps"
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
	app.Mount("/apps", apps.ApplicationRouting())

	app.Listen(port)
}
