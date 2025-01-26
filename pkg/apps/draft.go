package apps

import (
	mongodb "byteserver/pkg/mongo"
	_"byteserver/pkg/redis"
	"byteserver/pkg/utils"
	"log"
	"github.com/gofiber/fiber/v2"
)

/*
 Body: {
	cohort_id: "1",
	deadline: "YYYY-MM-DD HH:mm:ss",
	questions: [
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
	var params utils.EditDraft

	err := utils.Validate(c, &params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error with validating inputs",
			"error": err.Error(),
		})
	}

	if err = mongodb.EditDraft(params); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error Editing Draft",
			"error":   err.Error(),
		})
	}
	/*
	if err = redis.ClearCache(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error with Redis",
			"error":   err.Error(),
		})
	}*/
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully updated draft",
	})
}

/* Creates a new collection if it doesn't exist, and a draft document
 Example: /create-draft
 Body: {
	cohort_id: "1"
} */
func CreateDraft(c *fiber.Ctx) error {
	var params Cohort

	err := utils.Validate(c, &params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error with validating inputs",
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

func PublishDraft(c *fiber.Ctx) error {
	var params Cohort

	err := utils.Validate(c, &params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error with validating inputs",
			"error": err.Error(),
		})
	}

	err = mongodb.CreateTemplate("cohort-" + params.Cohort_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error Publishing Draft",
			"error":   err.Error(),
		})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully published draft",
	})
}