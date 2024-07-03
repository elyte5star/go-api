package routes

import (
	"fmt"

	"github.com/api/common/middleware"
	res "github.com/api/repository/response"
	"github.com/api/service"
	"github.com/gofiber/fiber/v2"
)

// @Summary Health Check
// @Description API status check
// @Tags API
// @Produce json
// @Success 200 {object} response.RequestResponse
// @Failure 500 {object} response.ErrorResponse
// @router /api/status [get]
func healthCheck(c *fiber.Ctx) error {
	response := res.NewResponse(c)
	response.Message = "Server is up and running"
	if err := c.Status(fiber.StatusOK).JSON(response); err != nil {
		return fmt.Errorf("error, Server is down, %w", err)
	}

	return nil
}

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(c *fiber.Ctx) error {
	response := res.NewErrorResponse()
	response.Message = "Sorry, endpoint is not found"
	response.Code = fiber.StatusNotFound
	return c.Status(fiber.StatusNotFound).JSON(response)
}

func MapRoutes(app *fiber.App, cfg *service.AppConfig) {

	//logger middleware
	//logger := cfg.Logger

	api := app.Group("api")
	serverStatus := api.Group("server")
	serverStatus.Get("/status", healthCheck)

	// JWT middleware
	jwt := middleware.NewAuthMiddleware(cfg.JwtSecretKey)

	authRoute := api.Group("auth")
	authRoute.Post("/login", cfg.Login)

	users := api.Group("users")
	users.Post("signup", cfg.CreateUser)
	authenticated := users.Use(jwt)
	authenticated.Get("", cfg.GetUsers)
	authenticated.Get("/:userid", cfg.GetUser)
	authenticated.Delete("/:userid", cfg.DeleteUser)
	authenticated.Put("/:userid", cfg.UpdateUser)

	productRoutes := app.Group("products")
	productRoutes.Get("/", cfg.GetAllProducts)
	productRoutes.Get("/:pid", cfg.GetSingleProduct)
	productRoutes.Delete("/:pid", jwt, cfg.DeleteProduct)
	productRoutes.Post("/create", cfg.CreateProduct)
	productRoutes.Post("/create/review", cfg.CreateProduct)

	// bookingRoutes := app.Group("/api/qbooking",jwt)
	// bookingRoutes.Post("/create")

	// jobRoute := app.Group("/api/job",jwt)
	// jobRoute.Get("/")
	// jobRoute.Get("/:jid")
	// jobRoute.Delete("/:jid")

	// NotFoundRoute func for describe 404 Error route.
	app.Use(NotFoundRoute)

}