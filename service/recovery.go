package service

import (
	"errors"
	"fmt"

	"github.com/api/repository/response"
	"github.com/gofiber/fiber/v2"
)

func (cfg *AppConfig) PanicRecovery(c *fiber.Ctx, err error) error {

	cfg.Logger.Error(err.Error())

	newErr := response.NewErrorResponse()
	// Status code defaults to 500
	statusCode := newErr.Code

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		statusCode = e.Code
		newErr.Code = statusCode
	}
	// Send custom error page
	err = c.Status(statusCode).SendFile(fmt.Sprintf("./%d.html", statusCode))

	if err != nil {
		// In case the SendFile fails
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}

	// Return from handler
	return nil

}
