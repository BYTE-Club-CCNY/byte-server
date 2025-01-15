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
				N_records:	c.QueryInt("n_records", 10),
				Page: 		c.QueryInt("page", 1),
			}

	var projects []database.GetProjects
	query := database.DB.Limit(params.N_records)

	err := query.Raw(database.GetProjectsQuery).Scan(&projects).Error
	// if params.Cohort != -1 {
	// 	query = query.Where("cohort_id = ?", params.Cohort)
	// }
	// if params.Name != "" {
	// 	query = query.Where("name = ?", params.Name)
	// }

	// err := query.Find(&projects)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(projects);
}