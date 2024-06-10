package routers

import (
	"fmt"
	"os"

	"github.com/api/common/config"
	"github.com/api/common/middleware"
	res "github.com/api/repository/response"
	"github.com/api/service"
	"github.com/gofiber/fiber/v2"
)

func healthCheck(c *fiber.Ctx) error {
	response := res.NewResponse(c)
	response.Message = "Server is up and running"
	if err := c.Status(fiber.StatusOK).JSON(response); err != nil {
		return fmt.Errorf("error, Server is down, %w", err)
	}

	return nil
}

func NotFoundRoute(c *fiber.Ctx) error {
	response := res.NewErrorResponse()
	response.Message = "Sorry, endpoint is not found"
	response.Code = fiber.StatusNotFound
	return c.Status(fiber.StatusNotFound).JSON(response)
}

func RouteSetup(app *fiber.App, cfg *config.AppConfig) {

	//logger middleware
	logger := cfg.Logger

	serverStatus := app.Group("/")
	serverStatus.Get("/status", healthCheck)

	//middleware
	// jwt := middleware.NewAuthMiddleware(util.JwtSecret)

	// productRoutes := app.Group("/api/products")
	// productRoutes.Get("/", service.GetAllProducts)
	// productRoutes.Get("/:pid", service.GetSingleProduct)
	// productRoutes.Delete("/:pid",jwt, service.DeleteProduct)

	userRoutes := app.Group("/api/users")
	// userRoutes.Get("/")
	userRoutes.Get("/:userid", service.GetUser)
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
		swaggerRoute := app.Group("/docs")
		swaggerRoute.Get("*", middleware.SwaggerHandler(cfg))
	} else {
		logger.Warn(fmt.Sprintf("Swagger file not found at %s, skipping redoc init", specFile))
	}
	// NotFoundRoute func for describe 404 Error route.
	app.Use(NotFoundRoute)

}
