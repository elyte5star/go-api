package main

import (
	"fmt"

	"github.com/api/common/config"
	"github.com/api/common/middleware"
	_ "github.com/api/docs"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func Handler(cfg *config.AppConfig) *fiber.App {

	appInfo := fmt.Sprintf("%s:%s", cfg.ServiceName, cfg.Version)

	//logger middleware
	logger := cfg.Logger

	// Fiber instance
	fb := fiber.New(fiber.Config{
		AppName:           appInfo,
		EnablePrintRoutes: cfg.Debug,
		ErrorHandler:      cfg.PanicRecovery,
	})

	//check if application meets requirments
	meetSysRequirment := util.SysRequirment(cfg)
	if !meetSysRequirment {
		fb.Shutdown()
	}

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

	// Add a Favicon middleware handler
	fb.Use(favicon.New(favicon.Config{
		File: "./docs/favicon.ico",
	}))

	// Recover middleware
	fb.Use(recover.New())

	//Set logging to DEBUG LEVEL in Development
	fb.Use(slogfiber.New(logger))

	DocRoute := fb.Group("/swagger")
	DocRoute.Get("*", middleware.SwaggerHandler(cfg))

	return fb
}
