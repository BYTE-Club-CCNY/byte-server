package main

import (
	"byteserver/pkg/apps"
	mongodb "byteserver/pkg/mongo"
	"byteserver/pkg/projects"
	"byteserver/pkg/redis"
	"byteserver/pkg/users"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {	
	port := ":3000"
	app := fiber.New()
	
	// database.InitDB()
	redis.InitRedis()
	mongodb.Connect()
	
	app.Use(cors.New())
	
	app.Use(func(c *fiber.Ctx) error {
    	fmt.Printf("%s request for %s\n", c.Method(), c.Path())
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		if c.Method() != "GET" {
			return c.Next()
		}

		key, _ := json.Marshal(c.AllParams())

		if value, err := redis.GetCache(string(key)); err == nil {
			var returnVal any;

			err := json.Unmarshal([]byte(value), &returnVal)
			if err != nil {
				return c.Next()
			}

			return c.Status(fiber.StatusOK).JSON(returnVal)
		}
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BYTE Server is running!")
	})
	app.Mount("/projects", projects.Projects())
	app.Mount("/apps", apps.App())
	app.Mount("/user", users.Users())

	app.Listen("localhost" + port)
}
