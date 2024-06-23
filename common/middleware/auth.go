package middleware

import (
	"github.com/api/repository/response"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	
)



// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(secret)},
		ContextKey:   "jwt",
		ErrorHandler: abortAuthenticationFailed,
	})
}

func abortAuthenticationFailed(c *fiber.Ctx, err error) error {
	newErr := response.NewErrorResponse()
	newErr.Message = err.Error()
	newErr.Code = fiber.StatusUnauthorized
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		newErr.Code = fiber.ErrBadRequest.Code
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Send back the Unauthorized message
	return c.Status(newErr.Code).JSON(newErr)

}
