package main

import (
	"fmt"
	"os"
	"time"
	"github.com/api/service"
	"github.com/api/util"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mvrilo/go-redoc"
	fiberredoc "github.com/mvrilo/go-redoc/fiber"
	slogfiber "github.com/samber/slog-fiber"
)

func SwaggerHandler(cfg *service.AppConfig) fiber.Handler {
	// Add the handler to serve the redoc
	swaggerConfig := swagger.Config{
		BasePath: "/api",
		FilePath: cfg.Doc,
		Path:     "docs",
		Title:    fmt.Sprintf("%s:%s Documentation", cfg.ServiceName, cfg.Version),
	}
	return swagger.New(swaggerConfig)

}

func DocumentationHandler(cfg *service.AppConfig) fiber.Handler {
	// Add the handler to serve the redoc
	doc := redoc.Redoc{
		Title:       fmt.Sprintf("%s:%s Documentation", cfg.ServiceName, cfg.Version),
		Description: "Documentation for Elyte-Realm API",
		SpecFile:    cfg.Doc, //
		SpecPath:    "/swagger.json",
		DocsPath:    "/api/docs",
	}
	return fiberredoc.New(doc)

}

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

	if _, err := os.Stat(cfg.Doc); err == nil {
		//fb.Use(DocumentationHandler(cfg))
		fb.Use(SwaggerHandler(cfg))
	} else {
		logger.Warn(fmt.Sprintf("Swagger file not found at %s, skipping redoc init", cfg.Doc))
	}

	//Add routes
	service.MapUrls(fb, cfg)

	return fb
}
