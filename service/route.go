package service

import (
	"fmt"
	"github.com/api/common/middleware"
	res "github.com/api/repository/response"
	"github.com/gofiber/fiber/v2"
	
)

// @tags App
// @router /api/status [get]
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

func MapUrls(app *fiber.App, cfg *AppConfig) {

	//logger middleware
	//logger := cfg.Logger

	serverStatus := app.Group("/api")
	serverStatus.Get("/status", healthCheck)

	//middleware
	jwt := middleware.NewAuthMiddleware(cfg.JwtSecretKey)
	// productRoutes := app.Group("/api/products")
	// productRoutes.Get("/", service.GetAllProducts)
	// productRoutes.Get("/:pid", service.GetSingleProduct)
	// productRoutes.Delete("/:pid",jwt, service.DeleteProduct)
	api := app.Group("api")
	users := api.Group("users")
	users.Post("create", cfg.CreateUser)
	authenticated := users.Use(jwt)
	authenticated.Get("/", cfg.GetUsers)
	authenticated.Get("/:userid", cfg.GetUser)

	// userRoutes.Delete("/:userid")

	authRoute := api.Group("auth")
	authRoute.Post("/login", cfg.Login)

	// bookingRoutes := app.Group("/api/qbooking",jwt)
	// bookingRoutes.Post("/create")

	// jobRoute := app.Group("/api/job",jwt)
	// jobRoute.Get("/")
	// jobRoute.Get("/:jid")
	// jobRoute.Delete("/:jid")


	// NotFoundRoute func for describe 404 Error route.
	app.Use(NotFoundRoute)

}
