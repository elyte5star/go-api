package service

import (
	"fmt"
	"time"

	"github.com/api/common/config"
	"github.com/api/common/database"
	"github.com/api/common/database/schema"
	"github.com/api/repository/response"
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
func GetUser(c *fiber.Ctx) error {
	var cfg config.AppConfig
	newErr := response.NewErrorResponse()
	userid, err := uuid.Parse(c.Params("userid"))
	if err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "invalid user id"
		cfg.Logger.Error(fmt.Sprintf("invalid user id: %s", err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	db, err := database.DbWithQueries(&cfg)
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
func CreateUser(c *fiber.Ctx) error {
	var cfg config.AppConfig
	// Get now time.
	now := time.Now().UTC()
	// Create new User struct
	user := &schema.User{}
	newErr := response.NewErrorResponse()
	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Create database connection.
	db, err := database.DbWithQueries(&cfg)
	if err != nil {
		newErr.Message = "Couldnt connect to DB!"
		cfg.Logger.Error(fmt.Sprintf("Couldnt connect to DB!: %s", err))
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}

	audit := &schema.AuditEntity{CreatedAt: now, LastModifiedAt: now, LastModifiedBy: "none", CreatedBy: user.Email}
	// Set initialized default data for user:
	user.Userid = uuid.New()
	user.AuditInfo = *audit
	user.Discount = 0.0
	user.Admin = false
	user.Enabled = true
	user.FailedAttempt = 0
	user.AccountNonLocked = false
	// Validate user fields.
	if err := cfg.Validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = fmt.Sprintf("Errors : %v", util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	if err := db.CreateUser(user); err != nil {
		newErr.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = user.Userid
	return c.Status(fiber.StatusOK).JSON(response)
}
