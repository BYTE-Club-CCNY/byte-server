package projects

import (
	"byteserver/pkg/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type GetProjectsBody struct {
	Cohort 		int
	Name 		string
	N_records 	int
	Page 		int
}


func Projects() *fiber.App {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
    	fmt.Printf("Request for %s\n", c.Path())
		return c.Next()
	})


	app.Get("/", sayHi)
	app.Get("/get", get)

	return app;
}

func printQueries(c *fiber.Ctx) {
	for key, value := range c.Queries() {
		fmt.Printf("%s: %s\t", key, value)
	}
	fmt.Print("\n")
}

func printParams(c *fiber.Ctx) {
	for key, value := range c.AllParams() {
		fmt.Printf("%s: %s\t", key, value)
	}
	fmt.Print("\n")
}

func sayHi(c *fiber.Ctx) error {
	return c.SendStatus(200);
}

func get(c *fiber.Ctx) error {
	printQueries(c)
	params := GetProjectsBody{
				Cohort: 	c.QueryInt("cohort", -1),
				Name:	   	c.Query("name"),
				Page: 		c.QueryInt("page", 1),
			}

	var projects []database.GetProjects
	query := database.DB.Scopes(database.Paginate(params.Page))


	query.Table("projects.project AS p").
		Select(`p.id, p.name, p.topic, p.short_desc, p.long_desc, 
				p.link, p.image, p.tech_stack,
				CONCAT(u1.first_name, ' ', u1.middle_name, ' ', u1.last_name) AS member1,
				CONCAT(u2.first_name, ' ', u2.middle_name, ' ', u2.last_name) AS member2,
				CONCAT(u3.first_name, ' ', u3.middle_name, ' ', u3.last_name) AS member3,
				CONCAT(u4.first_name, ' ', u4.middle_name, ' ', u4.last_name) AS member4,
				c.cohort_name AS cohort`).
		Joins("INNER JOIN projects.teams t ON p.id = t.id").
		Joins("LEFT JOIN users.people u1 ON t.member1 = u1.uid").
		Joins("LEFT JOIN users.people u2 ON t.member2 = u2.uid").
		Joins("LEFT JOIN users.people u3 ON t.member3 = u3.uid").
		Joins("LEFT JOIN users.people u4 ON t.member4 = u4.uid").
		Joins("INNER JOIN users.cohort c ON p.cohort_id = c.cohort_id")

	if params.Cohort != -1 {
		query = query.Where("p.cohort_id = ?", params.Cohort)
	}
	if params.Name != "" {
		query = query.Where("LOWER(p.name) = LOWER(?)", params.Name)
	}

	query.Scan(&projects)

	if query.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": query.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(projects);
}