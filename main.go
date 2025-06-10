package main

import (
	"byteserver/pkg/apps"
	"byteserver/pkg/database"
	"byteserver/pkg/projects"

	// "byteserver/pkg/redis"
	"byteserver/pkg/users"
	// "encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {	
	port := ":10000" // default for render
	app := fiber.New()
	
	database.InitDB()
	// redis.InitRedis()
	// mongodb.Connect()
	
	app.Use(cors.New())
	
	app.Use(func(c *fiber.Ctx) error {
    	fmt.Printf("%s request for %s\n", c.Method(), c.Path())
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BYTE Server is running!")
	})
	app.Mount("/projects", projects.Projects())
	app.Mount("/apps", apps.App())
	app.Mount("/user", users.Users())

	app.Listen("0.0.0.0" + port)
}
