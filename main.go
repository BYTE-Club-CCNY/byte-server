package main

import (
	"byteserver/pkg/apps"
	"byteserver/pkg/database"
	"byteserver/pkg/projects"
	"byteserver/pkg/redis"
	"byteserver/pkg/users"
	"byteserver/pkg/utils"

	// "byteserver/pkg/mongo"

	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {	
	port := ":3000"
	app := fiber.New()
	utils.IshmamLoadEnv()
	// database.InitDB()
	redis.InitRedis()
	// mongodb.Connect()
	
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

	app.Listen(port)
}
