package service

import (
	"time"
	"github.com/api/repository/request"
	"github.com/api/repository/response"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (cfg *AppConfig) LoginForm(c *fiber.Ctx) error {
	user := c.FormValue("username")
	pass := c.FormValue("password")

	secret := cfg.JwtSecretKey

	// Set expires minutes count for secret key from .env file.
	minutesCount := cfg.JwtExpireMinutesCount
	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func (cfg *AppConfig) Login(c *fiber.Ctx) error {

	newErr := response.NewErrorResponse()

	tokenReq := new(request.LoginRequest)

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&tokenReq); err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid JSON body"
		cfg.Logger.Error(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Validate Login fields.
	if err := cfg.Validate.Struct(tokenReq); err != nil {
		// Return, if some fields are not valid.
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = util.ValidatorErrors(err)
		cfg.Logger.Error(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		newErr.Message = "Couldnt connect to DB!"
		cfg.Logger.Error(newErr.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	user, err := db.FindByCredentials(tokenReq.Username)

	if err != nil {
		newErr.Message = "user with the given username is not found!"
		cfg.Logger.Error(newErr.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = user
	return c.Status(fiber.StatusOK).JSON(response)

}
