package main

import (
	"evermos/config"
	"evermos/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	app := fiber.New(fiber.Config{
		StrictRouting: false,
	})

	app.Static("/uploads", "./uploads")

	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend Evermos API is Running!")
	})

	app.Listen(":8080")

	app.Static("/uploads", "./uploads")
}
