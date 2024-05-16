package main

import (
	router "github.com/api/router"
	"github.com/api/util"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	
	util.SystemInfo()
	// Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "Elyte Realm v1.0.1",
	})

	app.Use(logger.New(util.RequestLogConfig))
	// Routes
	router.RouteSetup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// start server
	app.Listen(":8080")
}

