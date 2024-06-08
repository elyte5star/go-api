package middleware

import (
	"fmt"
	"os"

	"github.com/api/common/config"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(app *fiber.App, cfg *config.AppConfig) {
	// Add the handler to serve the redoc
	specFile := cfg.Doc
	if _, err := os.Stat(specFile); err == nil {
		swaggerRoute := app.Group("/swagger")
		swaggerConfig := swagger.Config{
			Title:    fmt.Sprintf("%s:%s Documentation", cfg.ServiceName, cfg.Version),
			FilePath: specFile,
		}
		swagger.New(swaggerConfig)
		swaggerRoute.Get("*", swagger.New(swaggerConfig))
	} else {
		cfg.Logger.Warn(fmt.Sprintf("Swagger file not found at %s, skipping redoc init", specFile))

	}

}
