package apps

import (
	mongodb "byteserver/pkg/mongo"
	"byteserver/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Retrieves all applications for a given collection
// Example Usage: /collection-data?collection=cohort-1
func GetCollectionData(c *fiber.Ctx) error {
	collectionName := c.Query("collection")
	
	if collectionName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required parameter: season",
		})
	}

	/*
	We should verify the token
	token := c.GetRespHeader("token", "invalid")
	if !verify(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token, You are not authorized.",
		})
	}
	*/

	data, err := mongodb.GetAllData(c.Context(), collectionName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch applications",
			"error":   err.Error(),
		})
	}

	c.Status(fiber.StatusOK).JSON(data)
	return nil
}

/* Creates a new collection if it doesn't exist
 Example: /new-collection
 Body: {
	cohort_id: "1"
} */
func NewCohort(c *fiber.Ctx) error {
	cohort := &Cohort{} 

	/*
	We should verify the token
	token := c.GetRespHeader("token", "invalid")
	if !verify(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token, You are not authorized.",
		})
	}
	*/

	err := utils.Validate(c, cohort)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := mongodb.CreateNewCohort(c.Context(), "cohort-" + cohort.Cohort_id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating collection",
			"error":   err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Successfully created %s collection", "cohort-" + cohort.Cohort_id),
	});
}	