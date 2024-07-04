package service

import (
	"fmt"
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
// @Tags User
// @Accept json
// @Produce json
// @Param userid path string true "userid"
// @Success 200 {object} response.RequestResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Security BearerAuth
// @Router /api/users/{userid} [get]
func (cfg *AppConfig) GetUser(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	userid, err := uuid.Parse(c.Params("userid"))
	if err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid userid"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	user, err := db.GetUserById(userid)
	if err != nil {
		newErr.Message = "User with userid is not found!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	result := &response.GetUserResponse{Userid: user.Userid,
		LastModifiedAt:   user.AuditInfo.LastModifiedAt,
		CreatedAt:        user.AuditInfo.CreatedAt,
		Username:         user.UserName,
		Email:            user.Email,
		AccountNonLocked: user.AccountNonLocked,
		Admin:            user.Admin,
		IsUsing2FA:       user.IsUsing2FA,
		Enabled:          user.Enabled,
		Telephone:        user.Telephone,
		LockTime:         user.LockTime,
	}
	response := response.NewResponse(c)
	response.Result = result
	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateUser func for updating a user by userid.
// @Description Update User.
// @Summary Update user
// @Tags User
// @Accept json
// @Produce json
// @Param userid path string true "userid"
// @Param modify_user body request.ModifyUser true "Modify User"
// @Success 201 {object} response.RequestResponse
// @Security BearerAuth
// @Router /api/users/{userid} [put]
func (cfg *AppConfig) UpdateUser(c *fiber.Ctx) error {
	// Get claims from JWT.
	data := cfg.JwtCredentials(c)
	loggedInUserid := data["userid"].(string)
	newErr := response.NewErrorResponse()
	userid, err := uuid.Parse(c.Params("userid"))
	if err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid userid"
		cfg.Logger.Error(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	modifyUser := new(request.ModifyUser)
	// Check, if received JSON data is valid.
	if err := c.BodyParser(modifyUser); err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid JSON body"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Validate modifyUser fields.
	if err := cfg.Validate.Struct(modifyUser); err != nil {
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
	foundUser, err := db.GetUserById(userid)
	if err != nil {
		newErr.Message = "User with the given userid is not found!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	foundUser.Telephone = modifyUser.Telephone
	foundUser.UserName = modifyUser.Username
	foundUser.AuditInfo.LastModifiedAt = util.TimeNow()
	foundUser.AuditInfo.LastModifiedBy = loggedInUserid 
	if err := db.UpdateUser(foundUser.Userid, &foundUser); err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Code = fiber.StatusCreated
	response.Result = fmt.Sprintf("User with ID : %v was updated.", userid)
	return c.Status(response.Code).JSON(response)

}

func (cfg *AppConfig) CreateUpdateUserAddress(userid uuid.UUID, address *request.CreateAddressReq) error {
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return err
	}
	userAddress := &schema.UserAddress{Userid: userid,
		FullName: address.FullName, StreetAddress: address.StreetAddress, Country: address.Country, State: address.State, Zip: address.Zip}
	if err := cfg.Validate.Struct(userAddress); err != nil {
		cfg.Logger.Error(util.ValidatorErrors(err))
		return err
	}
	if err := db.CreateUserAdress(userAddress); err != nil {
		cfg.Logger.Error(err.Error())
		return err
	}
	return nil

}

// CreateUser func creates a new user.
// @Description Create a new user.
// @Summary Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param create_user body request.CreateUserRequest true "Create User"
// @Success 200 {object} response.RequestResponse
// @Router /api/users/signup [post]
func (cfg *AppConfig) CreateUser(c *fiber.Ctx) error {

	newErr := response.NewErrorResponse()

	createUser := new(request.CreateUserRequest)

	// Check, if received JSON data is valid.
	if err := c.BodyParser(createUser); err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid JSON body"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Validate createUser fields.
	if err := cfg.Validate.Struct(createUser); err != nil {
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
	// Create new User struct
	user := new(schema.User)
	// Set initialized default data for user:
	user.Userid = util.Ident()
	user.UserName = createUser.Username
	user.SetPassword(createUser.Password)
	user.Email = createUser.Email
	user.LockTime = util.TimeThen()
	user.Telephone = createUser.Telephone
	audit := &schema.AuditEntity{CreatedAt: util.TimeNow(), LastModifiedAt: util.NullTime(), LastModifiedBy: "none", CreatedBy: user.Userid.String()}
	user.AuditInfo = *audit

	// Validate user fields.
	if err := cfg.Validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		cfg.Logger.Error(util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	if err := db.CreateUser(user); err != nil {
		newErr.Message = err.Error()
		if strings.Contains(err.Error(), "Error 1062") {
			newErr.Message = "Duplicate key: user already exist"
		}
		cfg.Logger.Error(err.Error())
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
// @Tags User
// @Accept json
// @Produce json
// @Failure 500 {object} response.ErrorResponse
// @Success 200 {array} response.RequestResponse
// @Security BearerAuth
// @Router /api/users [get]
func (cfg *AppConfig) GetUsers(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	//Get claims from JWT.
	data := cfg.JwtCredentials(c)
	isAdmin := data["isAdmin"].(bool)

	if !isAdmin {
		newErr.Message = "Admin rights needed"
		newErr.Code = fiber.StatusForbidden
		cfg.Logger.Warn(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	users, err := db.GetUsers()
	if err != nil {
		newErr.Message = "Users not found!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	// Define users variable.
	result := response.GetUsersResponse{}
	for _, user := range users {
		result.Users = append(result.Users, response.GetUserResponse{Userid: user.Userid,
			LastModifiedAt:   user.AuditInfo.LastModifiedAt,
			CreatedAt:        user.AuditInfo.CreatedAt,
			Username:         user.UserName,
			Email:            user.Email,
			AccountNonLocked: user.AccountNonLocked,
			Admin:            user.Admin,
			IsUsing2FA:       user.IsUsing2FA,
			Enabled:          user.Enabled,
			Telephone:        user.Telephone,
			LockTime:         user.LockTime,
		})
	}
	response := response.NewResponse(c)
	response.Result = result
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteUser func deletes user by a given userid.
// @Description Delete user by a given userid.
// @Summary Delete user by given userid
// @Tags User
// @Accept json
// @Produce json
// @Param userid path string true "userid"
// @Success 200 {object} response.RequestResponse
// @Security BearerAuth
// @Router /api/users/{userid} [delete]
func (cfg *AppConfig) DeleteUser(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	userid, err := uuid.Parse(c.Params("userid"))
	if err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid userid"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	db, err := DbWithQueries(cfg)
	if err != nil {
		newErr.Message = "Couldnt connect to DB!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	foundUser, err := db.GetUserById(userid)
	if err != nil {
		newErr.Message = "User with the given ID is not found!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	// Delete User by given user.
	if err := db.DeleteUser(foundUser.Userid); err != nil {
		newErr.Message = "Couldnt delete user!"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	response := response.NewResponse(c)
	return c.Status(response.Code).JSON(response)

}
