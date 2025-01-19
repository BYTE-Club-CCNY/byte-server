package apps

import (
	"byteserver/pkg/utils"
	"byteserver/pkg/redis"

	"github.com/gofiber/fiber/v2"
)


func EditCohortDraft(c *fiber.Ctx) error {
	var params EditCohortDraftBody

	err := utils.Validate(c, &params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	redis.ClearCache()
	return c.SendStatus(fiber.StatusOK)
}

/* Creates a new collection if it doesn't exist
 Example: /create-draft
 Body: {
	cohort_id: "1", 
	docType: "draft",
	<Optional> question-1: "question",
	<Optional> question-2: "question",
} */
func CreateDraft(c *fiber.Ctx) error {
	var params InitDraft

	err := utils.Validate(c, &params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func ViewDraft(c *fiber.Ctx) error {
	var params EditCohortDraftBody

	err := utils.Validate(c, &params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}