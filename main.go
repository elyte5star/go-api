package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// start server
	app.Listen(":8080")
}


