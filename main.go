package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello %s", c.Params("name"))
		return c.SendString(msg)
	})
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello %s, %s years old, %s", c.Params("name"), c.Params("age"), c.Params("gender"))
		return c.JSON(fiber.Map{"msg": msg})
		// return c.JSON(struct{"msg": msg})
	})
	app.Use(func(c *fiber.Ctx) error {
		basicauth.New(basicauth.Config{})
		return c.Next()
	})
	app.Listen(":3800")
}
