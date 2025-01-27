package apps

import (
	mongodb "byteserver/pkg/mongo"
	_ "byteserver/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

/*
{
	cohort_id: "cohort id (represented as a non-negative number)",
	user_id: "uuid of the user",
	"question": "answer"
	...
	"question": "answer"
}
*/
func SubmitApp(c *fiber.Ctx) error {
	var params bson.M

	if err := c.BodyParser(&params); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error Parsing Body",
			"error": err.Error(),
		})
	}

	if err := mongodb.UpdateOrInsertJSON(params, true); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error inserting/updating application",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully submitted application",
	})
}

func SaveApp(c *fiber.Ctx) error {
	var params bson.M

	if err := c.BodyParser(&params); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error Parsing Body",
			"error": err.Error(),
		})
	}

	if err := mongodb.UpdateOrInsertJSON(params, false); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error inserting/updating application",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully created/updated application",
	})
}

func ViewApps(c *fiber.Ctx) error {
	cohort_id := c.Query("cohort_id")
	pages, err := strconv.Atoi(c.Query("pages", "1"))
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Expected Int for pages",
		})
	}


	if cohort_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required parameter: cohort_id",
		})
	}

	collectionName := "cohort-" + cohort_id

	docs, err := mongodb.GetApps(collectionName, pages, limit) 
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error getting application documents",
			"error" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(docs)
}