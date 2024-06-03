package main

import (
	"fmt"
	slogfiber "github.com/samber/slog-fiber"
	"github.com/api/common/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/api/common/middleware"
)

func Handler(cfg *config.AppConfig) *fiber.App {

	appInfo := fmt.Sprintf("%s:%s", cfg.ServiceName, cfg.Version)

	// Fiber instance
	fb := fiber.New(fiber.Config{
		AppName:           appInfo,
		EnablePrintRoutes: cfg.Debug,
		ErrorHandler:      cfg.PanicRecovery,
	})

	// The index route is open
	fb.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "Ok",
		})
	})

	// Add a CORS middleware handler
	fb.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CorsOrigins,
	}))

	// Recover middleware
	fb.Use(recover.New())

	//logging middleware handler
	logger:= middleware.Logger()
	fb.Use(slogfiber.New(logger))



	return fb
}
