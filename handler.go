package main

import (
	"fmt"
	"time"

	"github.com/api/service"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func Handler(cfg *service.AppConfig) *fiber.App {

	appInfo := fmt.Sprintf("%s:%s", cfg.ServiceName, cfg.Version)

	//logger middleware
	logger := cfg.Logger

	// Fiber instance
	fb := fiber.New(fiber.Config{
		AppName:           appInfo,
		EnablePrintRoutes: cfg.Debug,
		ErrorHandler:      cfg.PanicRecovery,
		ReadTimeout:       time.Second * time.Duration(cfg.ReadTimeout),
	})

	//check if application meets requirments
	meetSysRequirment := util.SysRequirment(cfg.Logger)
	if !meetSysRequirment {
		fb.Shutdown()
	}

	// Add a CORS middleware handler
	fb.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     cfg.CorsOrigins,
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
	service.MapUrls(fb, cfg)

	return fb
}
