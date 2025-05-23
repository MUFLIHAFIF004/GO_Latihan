package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./views/index.html")
	})

	app.Get("/golang", func(c *fiber.Ctx) error {
		return c.SendFile("./views/golang.html")
	})
	app.Listen(":3000")
}
