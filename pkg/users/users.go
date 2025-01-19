package users

import (
	"byteserver/pkg/database"
	"byteserver/pkg/redis"
	schema "byteserver/pkg/schemas"
	"byteserver/pkg/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Users() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/add", add)

	return app
}

func add(c *fiber.Ctx) error {
	utils.PrintParams(c)
	var params AddUsersBody

	err := utils.Validate(c, &params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := new(schema.User)
	name := strings.Split(params.Name, " ")

	if middleName := strings.Join(name[1 : len(name) - 1], " "); strings.TrimSpace(middleName) != "" {
		user.MiddleName = middleName
	}

	user.FirstName = name[0]
	user.LastName = name[len(name) - 1]
	user.CunyEmail = params.CunyEmail
	user.Discord = params.Discord
	user.Emplid = params.Emplid
	user.PersonalEmail = params.PersonalEmail

	res := database.DB.Create(&user)

	if res.Error != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}
	
	redis.ClearCache()
	return c.Status(fiber.StatusOK).JSON(user)
}