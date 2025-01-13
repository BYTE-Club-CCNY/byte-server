package projects

import (
	"byteserver/pkg/database"
	schema "byteserver/pkg/schemas"

	"github.com/gofiber/fiber/v2"
)

type GetProjectsBody struct {
	Team 	string 	`json:"team" validate:"required"`
	Cohort 	string 	`json:"cohort" validate:"required"`
	Name 	string 	`json:"name" validate: "required"`
}

func Projects() *fiber.App {
	app := fiber.New()

	app.Get("/", sayHi);
	app.Get("/get", get)

	return app;
}

func sayHi(c *fiber.Ctx) error {
	return c.SendStatus(200);
}

func get(c *fiber.Ctx) error {
	params := GetProjectsBody{
			Team:   c.Query("team"),
			Cohort: c.Query("cohort"),
			Name:   c.Query("name"),
			}

	var projects []schema.Project
	query := database.DB.Limit(10)

	if params.Team != "" {
		query = query.Or("team = ?", params.Team)
	}
	if params.Cohort != "" {
		query = query.Or("cohort = ?", params.Cohort)
	}
	if params.Name != "" {
		query = query.Or("name = ?", params.Name)
	}

	result := query.Find(&projects)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Something went wrong internally",
		})
	}

	return c.Status(fiber.StatusOK).JSON(projects);

}