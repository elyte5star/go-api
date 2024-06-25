package middleware

import (
	"log"

	"github.com/api/repository/response"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler:   abortAuthenticationFailed,
		ContextKey:     "jwt",
		SuccessHandler: AuthSuccess,
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
	//increase failed attempt counter
	return c.Status(newErr.Code).JSON(newErr)

}

func AuthSuccess(c *fiber.Ctx) error {
	//reset failed attempt counter
	loggedInUser := c.Locals("jwt").(*jwt.Token)
	claims := loggedInUser.Claims.(jwt.MapClaims)
	userCredentials := claims["data"].(map[string]interface{})
	username := userCredentials["username"].(string)
	log.Printf("user '%s' accessing to '%s'", username, c.Request().URI().String())
	return c.Next()
}
