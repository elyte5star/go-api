package service

import (
	"fmt"
	"os"
	res "github.com/api/repository/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func healthCheck(c *fiber.Ctx) error {
	response := res.NewResponse(c)
	response.Message = "Server is up and running"
	if err := c.Status(fiber.StatusOK).JSON(response); err != nil {
		return fmt.Errorf("error, Server is down, %w", err)
	}

	return nil
}


func SwaggerHandler(cfg *AppConfig) fiber.Handler {
	// Add the handler to serve the redoc
	swaggerConfig := swagger.Config{
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		Title:        fmt.Sprintf("%s:%s Documentation", cfg.ServiceName, cfg.Version),
	}
	return swagger.New(swaggerConfig)

}
func NotFoundRoute(c *fiber.Ctx) error {
	response := res.NewErrorResponse()
	response.Message = "Sorry, endpoint is not found"
	response.Code = fiber.StatusNotFound
	return c.Status(fiber.StatusNotFound).JSON(response)
}

func RouteSetup(app *fiber.App, cfg *AppConfig) {

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
	userRoutes.Get("/",cfg.GetUsers)
	userRoutes.Get("/:userid", cfg.GetUser)
	userRoutes.Post("create", cfg.CreateUser)
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
		swaggerRoute.Get("*", SwaggerHandler(cfg))
	} else {
		logger.Warn(fmt.Sprintf("Swagger file not found at %s, skipping redoc init", specFile))
	}
	// NotFoundRoute func for describe 404 Error route.
	app.Use(NotFoundRoute)

}
