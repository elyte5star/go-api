package middleware

import (
	"fmt"
	"github.com/api/common/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerHandler(cfg *config.AppConfig) fiber.Handler {
	// Add the handler to serve the redoc
	swaggerConfig := swagger.Config{
		Title: fmt.Sprintf("%s:%s Documentation", cfg.ServiceName, cfg.Version),
	}
	return swagger.New(swaggerConfig)

}


