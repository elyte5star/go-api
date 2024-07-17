package service

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/api/repository/request"
	"github.com/api/repository/response"
	"github.com/api/repository/schema"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// const bearerPrefix = "Bearer "

// Login method for create a new bearer token.
// @Description Create a new bearer token.
// @Summary Create a new bearer token
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username/Email"
// @Param password formData string true "Password"
// @Success 200 {object} response.RequestResponse "OK"
// @Failure 400 {object} response.ErrorResponse{message=string,code=int} "BAD REQUEST"
// @Failure 404 {object} response.ErrorResponse{message=string,int} "NOT FOUND"
// @Failure 423 {object} response.ErrorResponse{message=string,code=int} "LOCKED"
// @Failure 503 {object} response.ErrorResponse{message=string,int} "SERVICE UNAVAILABLE"
// @Router /api/auth/form-login [post]
func (cfg *AppConfig) FormLogin(c *fiber.Ctx) error {

	newErr := response.NewErrorResponse()
	username := c.FormValue("username")
	password := c.FormValue("password")
	tokenReq := &request.LoginRequest{
		Username: username,
		Password: password,
	}
	// Validate Login fields.
	if err := cfg.Validate.Struct(tokenReq); err != nil {
		// Return, if some fields are not valid.
		newErr.Code = fiber.StatusBadRequest
		newErr.Message = fmt.Sprintf("Invalid Field(s) :%v", util.ValidatorErrors(err))
		cfg.Logger.Error(util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	user, err := schema.User{}, *new(error)
	if isEmail(tokenReq.Username) {
		user, err = db.FindByEmail(tokenReq.Username)
	} else {
		user, err = db.FindByUsername(tokenReq.Username)
	}
	if err != nil {
		newErr.Message = "User with the given username/email is not found!"
		newErr.Code = fiber.StatusNotFound
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	if !isEnabled(&user) || !isAccountNonLocked(&user) {
		newErr.Message = "Account not active or locked"
		newErr.Code = fiber.ErrLocked.Code
		return c.Status(newErr.Code).JSON(newErr)
	}
	if err = user.ComparePassword(tokenReq.Password); err != nil {
		newErr.Message = "Invalid password!"
		newErr.Code = fiber.StatusUnauthorized
		cfg.Logger.Error(err.Error())
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
	return c.Status(response.Code).JSON(response)

}

// Login method for create a new bearer token.
// @Description Create a new bearer token.
// @Summary Create a new bearer token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credential body request.LoginRequest true "Login data"
// @Success 200 {object} response.RequestResponse "OK"
// @Failure 400 {object} response.ErrorResponse{message=string,code=int} "BAD REQUEST"
// @Failure 404 {object} response.ErrorResponse{message=string,code=int} "NOT FOUND"
// @Failure 423 {object} response.ErrorResponse{message=string,code=int} "LOCKED"
// @Failure 503 {object} response.ErrorResponse{message=string,code=int} "SERVICE UNAVAILABLE"
// @Router /api/auth/login [post]
func (cfg *AppConfig) Login(c *fiber.Ctx) error {

	newErr := response.NewErrorResponse()

	tokenReq := new(request.LoginRequest)

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&tokenReq); err != nil {
		newErr.Code = fiber.StatusBadRequest
		newErr.Message = "Invalid JSON body"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Validate Login fields.
	if err := cfg.Validate.Struct(tokenReq); err != nil {
		// Return, if some fields are not valid.
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = fmt.Sprintf("Invalid Field(s) :%v", util.ValidatorErrors(err))
		cfg.Logger.Error(util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	user, err := schema.User{}, *new(error)
	if isEmail(tokenReq.Username) {
		user, err = db.FindByEmail(tokenReq.Username)
	} else {
		user, err = db.FindByUsername(tokenReq.Username)
	}
	if err != nil {
		newErr.Message = "User with the given username or email is not found!"
		newErr.Code = fiber.StatusNotFound
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	if !isEnabled(&user) || !isAccountNonLocked(&user) {
		newErr.Message = "Account not active or locked"
		newErr.Code = fiber.ErrLocked.Code
		return c.Status(newErr.Code).JSON(newErr)
	}
	if err = user.ComparePassword(tokenReq.Password); err != nil {
		newErr.Message = "Invalid password!"
		newErr.Code = fiber.StatusUnauthorized
		cfg.Logger.Error(err.Error())
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
	return c.Status(response.Code).JSON(response)

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
			Enabled:          user.Enabled,
			AccessToken:      token,
			TokenType:        "bearer",
		}, nil
	}
	return tokenResponse, err
}
func (cfg *AppConfig) GenerateJWT(user schema.User) (string, error) {

	principal := &request.UserCredentials{
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
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * time.Duration(cfg.JwtExpireMinutesCount)).Unix(),
		"data": principal,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(cfg.JwtSecretKey))

}
func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isAccountNonLocked(user *schema.User) bool {
	return user.AccountNonLocked
}

func isEnabled(user *schema.User) bool {
	return user.Enabled
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
