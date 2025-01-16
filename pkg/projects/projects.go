package projects

import (
	"byteserver/pkg/database"
	schema "byteserver/pkg/schemas"
	"byteserver/pkg/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

var validate = validator.New()

func Projects() *fiber.App {
	app := fiber.New()

	app.Get("/", sayHi)
	app.Get("/get", get)
	app.Post("/add", add)

	return app;
}

func parseArrays(array string) pq.StringArray {
	if array[0] != '[' || array[len(array) - 1] != ']' {
		panic("Malformed array")
	}

	// remove [] and "" 
	array = array[1 : len(array) - 1]
	var res []string = strings.Split(array, ",")
	for i := range res {
		res[i] = strings.TrimSpace(res[i])
	}

	return pq.StringArray(res)
}

func sayHi(c *fiber.Ctx) error {
	return c.SendStatus(200);
}

func get(c *fiber.Ctx) error {
	utils.PrintQueries(c)
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

func add(c *fiber.Ctx) error {
	utils.PrintParams(c)
	var params AddProjectsBody

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// get each person's UID
	var peopleNames []string
	members := []string{params.Member1, params.Member2, params.Member3, params.Member4}
	for i := range members {
		if members[i] != "" {
			peopleNames = append(peopleNames, strings.TrimSpace(members[i]))
		}
	}

	var peopleIDs []uuid.UUID;
	query := database.DB.Select("uid").Model(&schema.User{})
	err := query.Where("first_name IN (?)", peopleNames).Scan(&peopleIDs)
	
	if err.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error.Error())
	}

	if len(peopleIDs) != len(peopleNames) {
		return c.Status(fiber.ErrBadRequest.Code).SendString(`Some or all of these people do not exist`)
	}

	// create team & return team ID
	team := new(schema.Team)
	if len(peopleIDs) > 0 {
		team.Member1 = peopleIDs[0]
	}
	if len(peopleIDs) > 1 {
		team.Member2 = peopleIDs[1]
	}
	if len(peopleIDs) > 2 {
		team.Member3 = peopleIDs[2]
	} 
	if len(peopleIDs) > 3 {
		team.Member3 = peopleIDs[3]
	}

	res := database.DB.Create(&team)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.Error.Error())
	}

	// create project with team ID
	var project schema.Project
	project.Name = params.Name
	project.ID = team.ID
	project.Link = params.Link
	project.LongDesc = params.Long_Desc
	project.ShortDesc = params.Short_Desc
	project.Name = params.Name
	project.TechStack = parseArrays(params.Tech_Stack)
	project.Topic = parseArrays(params.Topic)
	project.CohortID = params.Cohort
	project.Image = params.Image

	database.DB.Create(project)
	

	return c.SendStatus(fiber.StatusOK)
}