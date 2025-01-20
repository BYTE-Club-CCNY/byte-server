package apps

import (
	mongodb "byteserver/pkg/mongo"
	"byteserver/pkg/redis"
	"byteserver/pkg/utils"
	"log"
	"github.com/gofiber/fiber/v2"
)

/*
 Body: {
	cohort_id: "1",
	deadline: "YYYY-MM-DD HH:mm:ss",
	<optional> questions: [
		{
			question: "string",
			type: "multiple-choice",
			options: ["string", "string", "string"]
		},
		{
			question: "string",
			type: "check-box",
			options: ["string", "string"]
		},
		{
			question: "string",
			type: "short-answer",
			options: []
		},
	]
} */
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
	cohort_id: "1"
} */
func CreateDraft(c *fiber.Ctx) error {
	var params Cohort

	err := utils.Validate(c, &params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	 
	err = mongodb.CreateDraft(params.Cohort_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error Creating Draft",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully created draft",
	})
}


// Example Usage: /view-draft?cohort_id=1
func ViewDraft(c *fiber.Ctx) error {
	cohort_id := c.Query("cohort_id")
	
	if cohort_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required parameter: cohort_id",
		})
	}

	log.Println(cohort_id)

	draft, err := mongodb.ViewDraft("cohort-" + cohort_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error Finding/Viewing Draft",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(draft)
}