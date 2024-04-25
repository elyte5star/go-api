package main

import "github.com/gofiber/fiber"

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// start server
	app.Listen(3000)
}

// Handler
func hello(c *fiber.Ctx) {
	c.send("Hello, world!")
}
