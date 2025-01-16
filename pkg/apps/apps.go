package apps

import (
	"github.com/gofiber/fiber/v2"
	"byteserver/pkg/mongo"
	"fmt"
)

func ApplicationRouting() *fiber.App {
	app := fiber.New()

	app.Get("/season-apps", GetSeasonApplications)
	app.Get("/new-season", CreateNewSeason)
	
	return app
}

// Retrieves all applications for a given season
func GetSeasonApplications(c *fiber.Ctx) error {
	season := c.Query("season")
	fmt.Println(season)

	/*
	We should verify the token
	token := c.GetRespHeader("token", "invalid")
	if !verify(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token, You are not authorized.",
		})
	}
	*/
	data := mongodb.GetAllApps(season)
	c.Status(fiber.StatusOK).JSON(data)
	return nil
}

// Creates a new collection for a new season if it doesn't exist
// Example Input: /new-season?season=spring-2025
func CreateNewSeason(c *fiber.Ctx) error {
	season := c.Query("season")

	/*
	We should verify the token
	token := c.GetRespHeader("token", "invalid")
	if !verify(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token, You are not authorized.",
		})
	}
	*/

	if err := mongodb.CreateNewCollection(&season); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating collection",
			"error": err.Error(),
		});
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully created collection",
	});
}	