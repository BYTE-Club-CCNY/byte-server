package main

import (
	"byteserver/pkg/database"
	"byteserver/pkg/projects"
	"byteserver/pkg/users"

	// "byteserver/pkg/mongo"
	"byteserver/pkg/applications"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {	
	port := ":3000"
	app := fiber.New()
	utils.IshmamLoadEnv()
	database.InitDB()
	redis.InitRedis()
	mongodb.Connect()
	
	app.Use(cors.New())
	
	app.Use(func(c *fiber.Ctx) error {
    	fmt.Printf("%s request for %s\n", c.Method(), c.Path())
		return c.Next()
	})

	/*
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() != "GET" {
			return c.Next()
		}

		key, _ := json.Marshal(c.AllParams())

		if value, err := redis.GetCache(string(key)); err == nil {
			return c.Status(fiber.StatusOK).SendString(value)
		}
		return c.Next()
	})*/

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BYTE Server is running!")
	})
	app.Mount("/projects", projects.Projects())
	app.Mount("/apps", apps.App())

	app.Listen(port)
}
