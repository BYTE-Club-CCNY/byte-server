package apps

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	_"byteserver/pkg/utils"
	mongodb "byteserver/pkg/mongo"
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