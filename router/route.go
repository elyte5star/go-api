package routers

import (
	"fmt"
	"os"
	"time"

	"github.com/api/common/config"
	"github.com/api/common/middleware"
	res "github.com/api/repository/response"
	"github.com/gofiber/fiber/v2"
)

func healthCheck(c *fiber.Ctx) error {

	res_ := res.RequestResponse{
		Path:      c.Route().Path,
		Message:   "Server is up and running",
		Success:   true,
		Code:      fiber.StatusOK,
		TimeStamp: time.Now().UTC(),
	}
	if err := c.Status(fiber.StatusOK).JSON(res_); err != nil {
		return fmt.Errorf("error, Server is down, %w", err)
	}

	return nil
}

func RouteSetup(app *fiber.App, cfg *config.AppConfig) {

	//logger middleware
	logger := cfg.Logger

	app.Get("/", healthCheck)

	//middleware
	// jwt := middleware.NewAuthMiddleware(util.JwtSecret)

	// productRoutes := app.Group("/api/products")
	// productRoutes.Get("/", service.GetAllProducts)
	// productRoutes.Get("/:pid", service.GetSingleProduct)
	// productRoutes.Delete("/:pid",jwt, service.DeleteProduct)

	// userRoutes := app.Group("/api/users",jwt)
	// userRoutes.Get("/")
	// userRoutes.Get("/:userid")
	// userRoutes.Delete("/:userid")

	// authRoute := app.Group("/api/auth")
	// authRoute.Post("/login")

	// bookingRoutes := app.Group("/api/qbooking",jwt)
	// bookingRoutes.Post("/create")

	// jobRoute := app.Group("/api/job",jwt)
	// jobRoute.Get("/")
	// jobRoute.Get("/:jid")
	// jobRoute.Delete("/:jid")

	specFile := cfg.Doc
	if _, err := os.Stat(specFile); err == nil {
		swaggerRoute := app.Group("/swagger")
		swaggerRoute.Get("*", middleware.SwaggerHandler(cfg))
	} else {
		logger.Warn(fmt.Sprintf("Swagger file not found at %s, skipping redoc init", specFile))
	}

}
