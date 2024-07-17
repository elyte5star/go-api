package routes

import (
	"encoding/json"
	"fmt"

	"github.com/api/common/middleware"
	res "github.com/api/repository/response"
	"github.com/api/service"
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
	response.Result = "Ok"
	if err := c.Status(response.Code).JSON(response); err != nil {
		return fmt.Errorf("error, Server is down, %w", err)
	}

	return nil
}

// NotFoundRoute func for describe 404 Error route.
func notFoundRoute(c *fiber.Ctx) error {
	response := res.NewErrorResponse()
	response.Message = "Sorry, endpoint is not found"
	response.Code = fiber.StatusNotFound
	return c.Status(response.Code).JSON(response)
}


// @Summary API Route information
// @Description API Route information
// @Tags API
// @Accept json
// @Produce json
// @Success 200 {object} response.RequestResponse "OK"
// @Failure 501 {object} response.ErrorResponse{message=string,code=int} "SERVICE UNAVAILABLE"
// @Security BearerAuth
// @Router /api/server/stack [get]
func routeStack(c *fiber.Ctx) error {
	data, err := json.MarshalIndent(c.App().Stack(), "", "  ")
	if err != nil {
		newErr := res.NewErrorResponse()
		return c.Status(newErr.Code).JSON(newErr)
	}
	response := res.NewResponse(c)
	response.Result = string(data)
	response.Message = "Checking your API Route information"
	return c.Status(response.Code).JSON(response)
}

func MapRoutes(app *fiber.App, cfg *service.AppConfig) {
	// JWT middleware
	jwt := middleware.NewAuthMiddleware(cfg.JwtSecretKey)
	api := app.Group("api")
	serverStatus := api.Group("server")
	serverStatus.Get("/status", jwt, healthCheck)
	serverStatus.Get("/stack", jwt, routeStack)

	authRoute := api.Group("auth")
	authRoute.Post("/login", cfg.Login)
	authRoute.Post("/form-login", cfg.FormLogin)

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
	productRoutes.Post("/create-many", jwt, cfg.CreateProducts)
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
