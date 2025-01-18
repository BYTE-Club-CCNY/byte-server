package main

import (
	"byteserver/pkg/database"
	"byteserver/pkg/projects"
	"byteserver/pkg/utils"

	//"byteserver/pkg/users"
	 "byteserver/pkg/mongo"
	"byteserver/pkg/apps"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {	
	port := ":3000"
	app := fiber.New()
	utils.IshmamLoadEnv()
	database.InitDB()
	mongodb.Connect()
	
	app.Use(func(c *fiber.Ctx) error {
    	fmt.Printf("%s request for %s\n", c.Method(), c.Path())
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BYTE Server is running!")
	})
	app.Mount("/projects", projects.Projects())
	app.Mount("/apps", apps.ApplicationRouting())

	app.Listen(port)
}
