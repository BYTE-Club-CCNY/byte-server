package projects

import "github.com/gofiber/fiber/v2"

func Projects() *fiber.App {
	app := fiber.New()

	app.Get("/", sayHi);

	return app;
}

func sayHi(c *fiber.Ctx) error {
	return c.SendStatus(200);
}