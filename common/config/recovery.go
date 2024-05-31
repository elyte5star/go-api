package config

import (
	"errors"
	"fmt"

	"github.com/api/common/response"
	"github.com/gofiber/fiber/v2"
)

func (cfg *AppConfig) PanicRecovery(c *fiber.Ctx, err error) error {

	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	cfg.Logger.Error(err.Error())
	// Send custom error page
	err = c.Status(code).SendFile(fmt.Sprintf("./%d.html", code))

	if err != nil {
		cfg.Logger.Error(err.Error())
		// In case the SendFile fails
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{Code: fiber.StatusInternalServerError, Cause: "Something went wrong"})
	}

	// Return from handler
	return nil

}
