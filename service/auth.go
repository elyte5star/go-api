package service

import (
	"time"

	"github.com/api/repository/request"
	"github.com/api/repository/response"
	"github.com/api/service/dbutils/schema"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserCredentials struct {
	Userid                  uuid.UUID `json:"userid"`
	Username                string    `json:"username"`
	Email                   string    `json:"email"`
	IsEnabled               bool      `json:"isEnabled "`
	IsAccountNonLocked      bool      `json:"isAccountNonLocked"`
	IsCredentialsNonExpired bool      `json:"isCredentialsNonExpired"`
	IsAdmin                 bool      `json:"isAdmin"`
	TokenId                 string    `json:"tokenId"`
}

//const bearerPrefix = "Bearer "

// Login method for create a new access token.
// @Description Create a new access token.
// @Summary Create a new access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credential body request.LoginRequest true "Login data"
// @Success 200 {object} response.RequestResponse
// @Router /api/auth/login [post]
func (cfg *AppConfig) Login(c *fiber.Ctx) error {
	// user := c.FormValue("username")
	// pass := c.FormValue("password")
	newErr := response.NewErrorResponse()

	tokenReq := new(request.LoginRequest)

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&tokenReq); err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid JSON body"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Validate Login fields.
	if err := cfg.Validate.Struct(tokenReq); err != nil {
		// Return, if some fields are not valid.
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid Field(s)"
		cfg.Logger.Error(util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	user, err := db.FindByCredentials(tokenReq.Username)
	if err != nil {
		newErr.Message = "User with the given username is not found!"
		newErr.Code = fiber.StatusNotFound
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	if err = user.ComparePassword(tokenReq.Password); err != nil {
		newErr.Message = "Invalid password!"
		newErr.Code = fiber.StatusUnauthorized
		cfg.Logger.Error(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	tokenResponse, err := cfg.GetTokenResponse(user)
	if err != nil {
		newErr.Message = "We could not log you in at this time, please try again later"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = tokenResponse
	return c.Status(fiber.StatusOK).JSON(response)

}
func (cfg *AppConfig) JwtCredentials(c *fiber.Ctx) map[string]interface{} {
	loggedInUser := c.Locals("jwt").(*jwt.Token)
	claims := loggedInUser.Claims.(jwt.MapClaims)
	userCredentials := claims["data"].(map[string]interface{})
	return userCredentials

}
func (cfg *AppConfig) GetTokenResponse(user schema.User) (response.TokenResponse, error) {
	tokenResponse := response.TokenResponse{}
	token, err := cfg.GenerateJWT(user)
	if err == nil {
		return response.TokenResponse{
			Userid:           user.Userid,
			Username:         user.UserName,
			Email:            user.Email,
			AccountNonLocked: user.AccountNonLocked,
			Admin:            user.Admin,
			AccessToken:      token,
			TokenType:        "bearer",
		}, nil
	}
	return tokenResponse, err
}
func (cfg *AppConfig) GenerateJWT(user schema.User) (string, error) {

	principal := &UserCredentials{
		Userid:                  user.Userid,
		Username:                user.UserName,
		Email:                   user.Email,
		IsEnabled:               user.Enabled,
		IsAccountNonLocked:      user.AccountNonLocked,
		IsAdmin:                 user.Admin,
		IsCredentialsNonExpired: true,
		TokenId:                 util.Ident().String(),
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"name": "Elyte Application",
		"exp":  time.Now().Add(time.Minute * time.Duration(cfg.JwtExpireMinutesCount)).Unix(),
		"data": principal,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(cfg.JwtSecretKey))

}

// func ExtractJwtCredentials(c *fiber.Ctx, cfg *AppConfig) (*UserCredentials, error) {
// 	userCredentials := UserCredentials{}
// 	token, err := verifyToken(c, cfg)
// 	if err != nil {
// 		return &userCredentials, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		userCredMap := claims["data"].(map[string]interface{})
// 		if err := m.Decode(userCredMap, &userCredentials); err != nil {
// 			cfg.Logger.Error(err.Error())
// 			return &userCredentials, err
// 		}
// 		return &userCredentials, nil
// 	}
// 	return &userCredentials, err
// }

// func GetTokenFromHeader(c *fiber.Ctx, cfg *AppConfig) string {
// 	authHeaderValue := c.Get("Authorization")
// 	if !strings.HasPrefix(authHeaderValue, bearerPrefix) {
// 		cfg.Logger.Error("No bearer token found in Authorization header")
// 		return ""
// 	}
// 	tokenString := strings.TrimPrefix(authHeaderValue, bearerPrefix)
// 	if len(tokenString) == 0 {
// 		cfg.Logger.Error("No bearer token found in Authorization header")
// 		return ""
// 	}
// 	return tokenString

// }

// func verifyToken(c *fiber.Ctx, cfg *AppConfig) (*jwt.Token, error) {
// 	tokenString := GetTokenFromHeader(c, cfg)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(cfg.JwtSecretKey), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return token, nil

// }
