package projects

import "github.com/gofiber/fiber/v2"

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
	// extract params
	params := new(GetProjectsBody);

	if err := c.BodyParser(params); err != nil {
		return err
	}

	// make request

	// return data
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"team": params.Team,
		"cohort": params.Cohort,
		"name" : params.Name,
	})

}