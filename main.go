package main

import (
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
)

func main() {
	
	util.SystemInfo()
	// Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "Elyte Realm v1.0.1",
	})

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// start server
	app.Listen(":8080")
}

