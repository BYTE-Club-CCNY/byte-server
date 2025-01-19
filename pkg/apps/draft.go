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