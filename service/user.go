package service

import (
	"fmt"

	"github.com/api/service/dbutils/schema"
	"github.com/api/service/request"
	"github.com/api/service/response"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetUser func gets User by given ID or 404 error.
// @Description Get User by given ID.
// @Summary get user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param userid path string true "User ID"
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
		cfg.Logger.Error(fmt.Sprintf("invalid user id: %s", err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	db, err := DbWithQueries(cfg)
	if err != nil {
		newErr.Message = "Couldnt connect to DB!"
		cfg.Logger.Error(fmt.Sprintf("Couldnt connect to DB!: %s", err))
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	user, err := db.GetUserById(userid)
	if err != nil {
		cfg.Logger.Error(fmt.Sprintf("Error while searching: %s", err))
		newErr.Message = "user with the given ID is not found!"
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = user
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateUser func creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
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
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		newErr.Message = "Couldnt connect to DB!"
		cfg.Logger.Error(fmt.Sprintf("Couldnt connect to DB!: %s", err))
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	// Create new User struct
	user := new(schema.User)
	// Set initialized default data for user:
	user.Userid = util.Ident()
	user.UserName = createUser.UserName
	user.Password = createUser.Password
	user.Email = createUser.Email
	user.Telephone = createUser.Telephone
	user.Discount = 0.0
	user.Admin = false
	user.Enabled = false
	user.FailedAttempt = 0
	user.AccountNonLocked = false
	audit := &schema.AuditEntity{CreatedAt: now, LastModifiedAt: util.NullTime(), LastModifiedBy: "none", CreatedBy: createUser.Email}
	user.AuditInfo = *audit
	// Validate user fields.
	if err := cfg.Validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = fmt.Sprintf("Errors : %v", util.ValidatorErrors(err))
		cfg.Logger.Error(fmt.Sprintf("%+v\n", newErr))
		return c.Status(newErr.Code).JSON(newErr)
	}
	if err := db.CreateUser(user); err != nil {
		newErr.Message = err.Error()
		cfg.Logger.Error(fmt.Sprintf("%+v\n", newErr))
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = user.Userid
	return c.Status(fiber.StatusOK).JSON(response)
}
