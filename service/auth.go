package service

import (
	//"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/api/repository/request"
	"github.com/api/repository/response"
	"github.com/api/service/dbutils/schema"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserPrincipal struct {
	Userid                  uuid.UUID `json:"userid"`
	Username                string    `json:"username"`
	Email                   string    `json:"email"`
	Exp                     time.Time `json:"exp,omitempty"`
	IsEnabled               bool      `json:"isEnabled "`
	IsAccountNonLocked      bool      `json:"isAccountNonLocked"`
	IsCredentialsNonExpired bool      `json:"isCredentialsNonExpired"`
	IsAdmin                 bool      `json:"isAdmin"`
	TokenId                 string    `json:"tokenId"`
}

const bearerPrefix = "Bearer "

func (cfg *AppConfig) Login(c *fiber.Ctx) error {
	// user := c.FormValue("username")
	// pass := c.FormValue("password")
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
	if err = user.ComparePassword(tokenReq.Password); err != nil {
		newErr.Message = "Invalid password!"
		newErr.Code = fiber.StatusUnauthorized
		cfg.Logger.Error(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	tokenResponse, err := GenerateJWT(user, cfg)
	if err != nil {
		newErr.Message = "We could not log you in at this time, please try again later"
		return c.JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = tokenResponse
	return c.Status(fiber.StatusOK).JSON(response)

}

func GenerateJWT(user schema.User, cfg *AppConfig) (response.TokenResponse, error) {

	principal := &UserPrincipal{
		Userid:                  user.Userid,
		Username:                user.UserName,
		Email:                   user.Email,
		IsEnabled:               user.Enabled,
		IsAccountNonLocked:      user.AccountNonLocked,
		IsAdmin:                 user.Admin,
		IsCredentialsNonExpired: true,
		TokenId:                 util.Ident().String(),
	}
	//credentials, err := json.Marshal(principal)
	// if err != nil {
	// 	cfg.Logger.Error(err.Error())
	// 	return response.TokenResponse{}, err
	// }
	// Create the Claims
	claims := jwt.MapClaims{
		"name": "Elyte Application",
		"exp":  time.Now().Add(time.Minute * time.Duration(cfg.JwtExpireMinutesCount)).Unix(),
		"data": principal,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(cfg.JwtSecretKey))
	return response.TokenResponse{
		Userid:           principal.Userid,
		Username:         principal.Username,
		Email:            principal.Email,
		AccountNonLocked: principal.IsAccountNonLocked,
		Admin:            principal.IsAdmin,
		AccessToken:      t,
		TokenType:        "bearer",
	}, err

}

func ExtractJwtCredentials(c *fiber.Ctx, cfg *AppConfig) (string, error) {
	token, err := verifyToken(c, cfg)
	if err != nil {
		return "", err
	}
	//principal := new(UserPrincipal)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println(claims["data"])
		return claims["data"].(string), nil
	}
	return "", err

}

func GetTokenFromHeader(c *fiber.Ctx, cfg *AppConfig) string {
	authHeaderValue := c.Get("Authorization")
	if !strings.HasPrefix(authHeaderValue, bearerPrefix) {
		cfg.Logger.Error("No bearer token found in Authorization header")
		return ""
	}
	tokenString := strings.TrimPrefix(authHeaderValue, bearerPrefix)
	if len(tokenString) == 0 {
		cfg.Logger.Error("No bearer token found in Authorization header")
		return ""
	}
	return tokenString

}

func verifyToken(c *fiber.Ctx, cfg *AppConfig) (*jwt.Token, error) {
	tokenString := GetTokenFromHeader(c, cfg)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil

}
