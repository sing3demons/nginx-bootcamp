package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello fiber")
	})

	app.Get("/v1", func(c *fiber.Ctx) error {
		return c.JSON("hello fiber v1")
	})
	app.Listen(":3000")
}
