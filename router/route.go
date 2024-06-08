package routers

import (
	"github.com/api/common/config"
	"github.com/gofiber/fiber/v2"
)

func RouteSetup(app *fiber.App,cfg *config.AppConfig) {

	
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
}
