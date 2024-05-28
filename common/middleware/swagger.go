package middleware

import (
	"fmt"
	"os"

	"github.com/api/common/config"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func SwaggerCfg(cfg config.AppConfig) fiber.Handler {
	// Add the handler to serve the redoc
	specFile := "./docs/swagger.json"
	if _, err := os.Stat(specFile); err == nil {
		swaggerConfig := swagger.Config{
			FilePath: "./docs/swagger.json",
			Title:    fmt.Sprintf("%s:%s Documentation", cfg.ServiceName, cfg.Version),
		}
		return swagger.New(swaggerConfig)
	}
	cfg.Logger.Warn(fmt.Sprintf("Swagger file not found at %s, skipping redoc init", specFile))
	return nil

}
