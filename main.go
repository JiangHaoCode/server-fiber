package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello %s", c.Params("name"))
		return c.SendString(msg)
	})
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello %s, %s years old, %s", c.Params("name"), c.Params("age"), c.Params("gender"))
		return c.SendString(msg)
	})
	app.Listen(":3800")
}
