package apps

import (
	mongodb "byteserver/pkg/mongo"

	"github.com/gofiber/fiber/v2"
)

// Retrieves data of the template document which has questions, and deadline of 
// a given cohort
func ViewTemplate(c *fiber.Ctx) error {
	cohort_id := c.Query("cohort_id")

	if cohort_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required parameter: cohort_id",
		})
	}

	template, err := mongodb.ViewTemplate("cohort-" + cohort_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error Finding/Viewing Template",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(template)


}