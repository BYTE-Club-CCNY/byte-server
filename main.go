package main

import (
	"byteserver/pkg/database"
	"byteserver/pkg/projects"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {	
	port := ":3000"
	app := fiber.New()
	database.InitDB()
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://www.byteccny.com, http://www.byteccny.com, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BYTE Server is running!")
	})
	app.Mount("/projects", projects.Projects())

	app.Listen(port)
}
