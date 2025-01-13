package projects

import (
	"byteserver/pkg/database"
	schema "byteserver/pkg/schemas"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type GetProjectsBody struct {
	Cohort 	int
	Name 	string
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
			Cohort: c.QueryInt("cohort"),
			Name:   c.Query("name"),
			}

	var projects []schema.Project
	var results []database.GetProjects

	query := database.DB.Limit(10)

	if params.Cohort != 0 { // horrible edge case!!!!
		query = query.Where("cohort_id = ?", params.Cohort)
	}
	if params.Name != "" {
		query = query.Where("name = ?", params.Name)
	}

	err := query.Find(&projects)
	if err.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Something went wrong internally",
		})
	}

	for i := range projects {
		var cohort schema.Cohort
		var team schema.Team
		// get cohort name
		var newProject database.GetProjects
		cohortQuery := database.DB.Where("cohort_id = ?", projects[i].CohortID)
		_ = cohortQuery.Find(&cohort)
		newProject.Cohort = cohort.CohortName

		teamQuery := database.DB.Where("id = ?", projects[i].ID)
		_ = teamQuery.Find(&team)
		
		newProject.Names = getNames([]string{team.Member1.String(), 
			team.Member2.String(), team.Member3.String(), team.Member4.String()})
		newProject.ID = projects[i].ID.String()
		newProject.Image = projects[i].Image
		newProject.Link = projects[i].Link
		newProject.LongDesc = projects[i].LongDesc
		newProject.ShortDesc = projects[i].ShortDesc
		newProject.ProjectName = projects[i].Name
		newProject.TechStack = projects[i].TechStack
		newProject.Topic = projects[i].Topic

		// get team members names
		results = append(results, newProject)
	}

	return c.Status(fiber.StatusOK).JSON(results);
}

func getNames(names []string) []string {
	var res []string
	for i := range names {
		var people schema.User
		query := database.DB.Where("uid = ?", names[i])
		_ = query.Find(&people)

		if people.MiddleName == "" {
			res = append(res, fmt.Sprintf("%s %s",
				people.FirstName, people.LastName))
		} else {
			res = append(res, fmt.Sprintf("%s %s %s",
				people.FirstName, people.MiddleName, people.LastName))
		}
	}
	return res
}