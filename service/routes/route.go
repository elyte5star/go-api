package routes

import (
	"encoding/json"
	"fmt"

	"github.com/api/common/middleware"
	res "github.com/api/repository/response"
	"github.com/api/service"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
)

// @Summary Health Check
// @Description API status check
// @Tags API
// @Accept json
// @Produce json
// @Success 200 {object} response.RequestResponse
// @Failure 500 {object} response.ErrorResponse
// @Security BearerAuth
// @Router /api/server/status [get]
func healthCheck(c *fiber.Ctx) error {
	response := res.NewResponse(c)
	response.Message = "Server is up and running"
	if err := c.Status(fiber.StatusOK).JSON(response); err != nil {
		return fmt.Errorf("error, Server is down, %w", err)
	}

	return nil
}

// NotFoundRoute func for describe 404 Error route.
func notFoundRoute(c *fiber.Ctx) error {
	response := res.NewErrorResponse()
	response.Message = "Sorry, endpoint is not found"
	response.Code = fiber.StatusNotFound
	return c.Status(fiber.StatusNotFound).JSON(response)
}

func RouteStack(app *fiber.App) string {
	defer util.TimeElapsed(util.TimeNow(), "Checking your API information")
	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	return string(data)
}

func MapRoutes(app *fiber.App, cfg *service.AppConfig) {

	//logger middleware
	//logger := cfg.Logger

	// JWT middleware
	jwt := middleware.NewAuthMiddleware(cfg.JwtSecretKey)
	api := app.Group("api")
	serverStatus := api.Group("server")
	serverStatus.Get("/status", jwt, healthCheck)

	authRoute := api.Group("auth")
	authRoute.Post("/login", cfg.Login)

	users := api.Group("users")
	users.Post("signup", cfg.CreateUser)
	authenticated := users.Use(jwt)
	authenticated.Get("", cfg.GetUsers)
	authenticated.Get("/:userid", cfg.GetUser)
	authenticated.Get("/:userid/address", cfg.GetAddressByUserid)
	authenticated.Delete("/:userid", cfg.DeleteUser)
	authenticated.Put("/:userid", cfg.UpdateUser)

	productRoutes := api.Group("products")
	productRoutes.Get("", cfg.GetAllProducts)
	productRoutes.Get("/:pid", cfg.GetSingleProduct)
	productRoutes.Get("/:pid/reviews", cfg.GetProductReviewsByPid)
	productRoutes.Delete("/:pid", jwt, cfg.DeleteProduct)
	productRoutes.Post("/create", jwt, cfg.CreateProduct)
	productRoutes.Post("/create/review", cfg.CreateReview)

	// bookingRoutes := app.Group("/api/qbooking",jwt)
	// bookingRoutes.Post("/create")

	// jobRoute := app.Group("/api/job",jwt)
	// jobRoute.Get("/")
	// jobRoute.Get("/:jid")
	// jobRoute.Delete("/:jid")

	// NotFoundRoute func for describe 404 Error route.
	app.Use(notFoundRoute)

}
