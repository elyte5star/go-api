package service

import (
	"errors"
	"fmt"

	"github.com/api/service/response"
	"github.com/gofiber/fiber/v2"
)

func (cfg *AppConfig) PanicRecovery(c *fiber.Ctx, err error) error {

	// Status code defaults to 500
	statusCode := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		statusCode = e.Code
	}
	cfg.Logger.Error(err.Error())
	// Send custom error page
	err = c.Status(statusCode).SendFile(fmt.Sprintf("./%d.html", statusCode))

	if err != nil {
		cfg.Logger.Error(err.Error())
		// In case the SendFile fails
		return c.Status(fiber.StatusInternalServerError).JSON(response.NewErrorResponse())
	}

	// Return from handler
	return nil

}
