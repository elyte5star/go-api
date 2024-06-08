package main

import (
	"fmt"
	"github.com/api/common/config"
	routers "github.com/api/router"
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

	// Add the request logging middleware handler to all service routes
	fb.Use(slogfiber.New(logger))
	
	//Add routes
	routers.RouteSetup(fb, cfg)

	return fb
}
