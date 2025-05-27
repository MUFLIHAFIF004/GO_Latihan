package main

import (
	"fmt"

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
			Name:   "Zakaruna",
			Tempat: "Bandung",
			Age:    21,
		})
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		var user User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
		}

		fmt.Println("Data diterima:", user) // Debug log

		return c.JSON(fiber.Map{
			"message": "User berhasil dibuat",
			"data":    user,
		})
	})

	app.Listen(":3000")
}
