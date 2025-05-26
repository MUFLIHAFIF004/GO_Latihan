package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name   string `json:"name"`
	Tempat string `json:"tempat"`
	Age    int    `json:"age"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(User{
			Name:   "Afif",
			Tempat: "Bandung",
			Age:    21,
		})
	})

	app.Listen(":3000")
}
