package service

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)


type TokenResponse struct {
	UserId      uuid.UUID `json:"userid"`
	UserName    string    `json:"username"`
	Email       string    `json:"email"`
	Enabled     bool      `json:"enabled"`
	Admin       bool      `json:"admin"`
	AccessToken string    `json:"accessToken"`
	TokenType   string    `json:"tokenType"`
}

func Login(c *fiber.Ctx) error {
	user := c.FormValue("username")
	pass := c.FormValue("password")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
