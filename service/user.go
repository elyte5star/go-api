package service

import (
	"strings"

	"github.com/api/repository/request"
	"github.com/api/repository/response"
	"github.com/api/service/dbutils/schema"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetUser func gets User by given ID or 404 error.
// @Description Get User by given ID.
// @Summary Get user by given userid
// @Tags Users
// @Accept json
// @Produce json
// @Param userid path string true "userid"
// @Success 200 {object} response.RequestResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users/{userid} [get]
func (cfg *AppConfig) GetUser(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	userid, err := uuid.Parse(c.Params("userid"))
	if err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "invalid user id"
		cfg.Logger.Error(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	db, err := DbWithQueries(cfg)
	if err != nil {
		newErr.Message = "Couldnt connect to DB!"
		cfg.Logger.Error(newErr.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	user, err := db.GetUserById(userid)
	if err != nil {
		newErr.Message = "user with the given ID is not found!"
		cfg.Logger.Error(newErr.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = user
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateUser func creates a new user.
// @Description Create a new user.
// @Summary Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Param confirmPassword body string true "ConfirmPassword"
// @Param email body string true "Email"
// @Param telephone body string true "telephone"
// @Success 200 {object} response.RequestResponse
// @Router /api/users/create [post]
func (cfg *AppConfig) CreateUser(c *fiber.Ctx) error {
	// Get now time.
	now := util.TimeNow()

	newErr := response.NewErrorResponse()

	createUser := new(request.CreateUserRequest)

	// Check, if received JSON data is valid.
	if err := c.BodyParser(createUser); err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = err.Error()
		cfg.Logger.Error(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Validate createUser fields.
	if err := cfg.Validate.Struct(createUser); err != nil {
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
	// Create new User struct
	user := new(schema.User)
	// Set initialized default data for user:
	user.Userid = util.Ident()
	user.UserName = createUser.UserName
	user.SetPassword(createUser.Password)
	user.Email = createUser.Email
	user.LockTime = util.TimeThen()
	user.Telephone = createUser.Telephone
	audit := &schema.AuditEntity{CreatedAt: now, LastModifiedAt: util.NullTime(), LastModifiedBy: "none", CreatedBy: user.Userid.String()}
	user.AuditInfo = *audit

	// Validate user fields.
	if err := cfg.Validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = util.ValidatorErrors(err)
		cfg.Logger.Error(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	if err := db.CreateUser(user); err != nil {
		newErr.Message = err.Error()
		if strings.Contains(err.Error(), "Error 1062") {
			newErr.Message = "Duplicate key: user alreay exist"
		}
		cfg.Logger.Error(newErr.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	//fmt.Printf("%+v\n", user)
	response := response.NewResponse(c)
	response.Result = user.Userid
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetUsers method for getting all existing users.
// @Description Get all existing users.
// @Summary Get all existing users
// @Tags Users
// @Accept json
// @Produce json
// @Failure 500 {object} response.ErrorResponse	
// @Success 200 {array} response.RequestResponse
// @Router /api/users/ [get]
func (cfg *AppConfig) GetUsers(c *fiber.Ctx) error {

	newErr := response.NewErrorResponse()
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		newErr.Message = "Couldnt connect to DB!"
		cfg.Logger.Error(newErr.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	users, err := db.GetUsers()
	if err != nil {
		newErr.Message = "users not found!"
		cfg.Logger.Error(newErr.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = users
	return c.Status(fiber.StatusOK).JSON(response)
}
