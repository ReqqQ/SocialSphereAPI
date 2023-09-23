package ui

import (
	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
}

func (co Controllers) InitGetRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Get("/user/1", func(c *fiber.Ctx) error {
		return c.JSON("test")
	})
}
