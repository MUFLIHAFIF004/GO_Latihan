package main

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./views/index.html")
	})

	app.Get("/golang", func(c *fiber.Ctx) error {
		return c.SendFile("./views/golang.html")
	})
}
