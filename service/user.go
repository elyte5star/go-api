package service

import (
	"fmt"
	"strings"

	"github.com/api/repository/request"
	"github.com/api/repository/response"
	"github.com/api/repository/schema"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateUser func creates a new user.
// @Description Create a new user.
// @Summary Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param create_user body request.CreateUserRequest true "Create User"
// @Success 201 {object} response.RequestResponse "CREATED"
// @Failure 400 {object} response.ErrorResponse{message=string,code=int} "BAD REQUEST"
// @Failure 409 {object} response.ErrorResponse{message=string,code=int} "CONFLICT"
// @Failure 501 {object} response.ErrorResponse{message=string,int} "SERVICE UNAVAILABLE"
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

	discount := 0.0
	if createUser.Discount == nil {
		createUser.Discount = &discount
	}
	// Create new User struct
	user := new(schema.User)
	// Set initialized default data for user:
	user.Userid = util.Ident()
	user.Username = createUser.Username
	user.SetPassword(createUser.Password)
	user.Email = createUser.Email
	user.Admin = false
	user.IsUsing2FA = false
	user.AccountNonLocked = true
	user.Enabled = false
	user.FailedAttempt = 0
	user.Discount = *createUser.Discount
	user.Telephone = createUser.Telephone
	audit := &schema.AuditEntity{CreatedAt: util.TimeNow(), LastModifiedBy: "none", CreatedBy: user.Userid.String()}
	user.AuditInfo = *audit

	// Validate user fields.
	if err := cfg.Validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		cfg.Logger.Error(util.ValidatorErrors(err))
		newErr.Message = fmt.Sprintf("Invalid Field(s) :%v", util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error("Couldnt connect to DB: " + err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	if err := db.CreateUser(user); err != nil {
		newErr.Message = err.Error()
		if strings.Contains(err.Error(), "Error 1062") {
			newErr.Message = "User with email or username or telephone already exist"
			newErr.Code = fiber.ErrConflict.Code
		}
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	//fmt.Printf("%+v\n", user)
	response := response.NewResponse(c)
	response.Result = user.Userid
	response.Code = fiber.StatusCreated
	return c.Status(response.Code).JSON(response)
}

// GetUser func gets User by given ID or 404 error.
// @Description Get User by given ID.
// @Summary Get user by given userid
// @Tags User
// @Accept json
// @Produce json
// @Param userid path string true "userid"
// @Success 200 {object} response.RequestResponse "OK"
// @Failure 400 {object} response.ErrorResponse{message=string,code=int} "BAD REQUEST"
// @Failure 404 {object} response.ErrorResponse{message=string,int} "NOT FOUND"
// @Failure 503 {object} response.ErrorResponse{message=string,int} "SERVICE UNAVAILABLE"
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
		cfg.Logger.Error("Couldnt connect to DB: " + err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	user, address, err := db.GetUserById(userid)
	if err != nil {
		newErr.Message = "User with userid is not found!"
		cfg.Logger.Error(err.Error())
		newErr.Code = fiber.StatusNotFound
		return c.Status(newErr.Code).JSON(newErr)
	}
	result := &response.GetUserResponse{Userid: user.Userid,
		LastModifiedAt:   user.AuditInfo.LastModifiedAt,
		CreatedAt:        user.AuditInfo.CreatedAt,
		Username:         user.Username,
		Email:            user.Email,
		AccountNonLocked: user.AccountNonLocked,
		Admin:            user.Admin,
		IsUsing2FA:       user.IsUsing2FA,
		Enabled:          user.Enabled,
		Telephone:        user.Telephone,
		LockTime:         user.LockTime,
		Address:          address,
	}
	response := response.NewResponse(c)
	response.Result = result
	return c.Status(response.Code).JSON(response)
}

// UpdateUser func for updating a user by userid.
// @Description Update User.
// @Summary Update user
// @Tags User
// @Accept json
// @Produce json
// @Param userid path string true "userid"
// @Param modify_user body request.ModifyUser true "Modify User"
// @Success 201 {object} response.RequestResponse "CREATED"
// @Failure 400 {object} response.ErrorResponse{message=string,code=int} "BAD REQUEST"
// @Failure 404 {object} response.ErrorResponse{message=string,int} "NOT FOUND"
// @Failure 409 {object} response.ErrorResponse{message=string,code=int} "CONFLICT"
// @Failure 501 {object} response.ErrorResponse{message=string,int} "SERVICE UNAVAILABLE"
// @Security BearerAuth
// @Router /api/users/{userid} [put]
func (cfg *AppConfig) UpdateUser(c *fiber.Ctx) error {
	// Get claims from JWT.
	data := cfg.JwtCredentials(c)

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
		newErr.Message = fmt.Sprintf("Invalid Field(s) :%v", util.ValidatorErrors(err))
		cfg.Logger.Error(newErr.Message)
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error("Couldnt connect to DB: " + err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	foundUser, address, err := db.GetUserById(userid)
	if err != nil {
		newErr.Message = "User with the given userid is not found!"
		cfg.Logger.Error(err.Error())
		newErr.Code = fiber.StatusNotFound
		return c.Status(newErr.Code).JSON(newErr)
	}
	modifyAddress := modifyUser.Address
	if modifyAddress != nil {
		newAddress := &schema.UserAddress{Userid: foundUser.Userid, FullName: modifyAddress.FullName,
			StreetAddress: modifyAddress.StreetAddress, Country: modifyAddress.Country,
			State: modifyAddress.State, Zip: modifyAddress.Zip}
		// Validate User Address fields.
		if err := cfg.Validate.Struct(newAddress); err != nil {
			// Return, if some fields are not valid.
			newErr.Code = fiber.ErrBadRequest.Code
			newErr.Message = fmt.Sprintf("Invalid Field(s) :%v", util.ValidatorErrors(err))
			cfg.Logger.Error(newErr.Message)
			return c.Status(newErr.Code).JSON(newErr)
		}
		if address.FullName == nil {
			if err := db.CreateUserAdress(newAddress); err != nil {
				cfg.Logger.Error(err.Error())
				return c.Status(newErr.Code).JSON(newErr)
			}
			cfg.Logger.Info("User Address was Created")
		} else {
			if err := db.UpdateUserAdress(foundUser.Userid, newAddress); err != nil {
				cfg.Logger.Error(err.Error())
				return c.Status(newErr.Code).JSON(newErr)
			}
			cfg.Logger.Info("User Address was Updated")
		}
		//announce change of address
	}
	foundUser.Telephone = modifyUser.Telephone
	foundUser.Username = modifyUser.Username
	now := util.TimeNow()
	foundUser.AuditInfo.LastModifiedAt = &now
	foundUser.AuditInfo.LastModifiedBy = data["username"].(string)
	if err := db.UpdateUser(foundUser.Userid, &foundUser); err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Code = fiber.StatusCreated
	response.Result = fmt.Sprintf("User with ID : %v was updated.", userid)
	return c.Status(response.Code).JSON(response)

}

// GetAddressByUserid func for getting a user's address by userid.
// @Description Ger User Address.
// @Summary Ger User Address
// @Tags User
// @Accept json
// @Produce json
// @Param userid path string true "userid"
// @Success 200 {object} response.RequestResponse "OK"
// @Failure 400 {object} response.ErrorResponse{message=string,code=int} "BAD REQUEST"
// @Failure 404 {object} response.ErrorResponse{message=string,int} "NOT FOUND"
// @Failure 503 {object} response.ErrorResponse{message=string,int} "SERVICE UNAVAILABLE"
// @Security BearerAuth
// @Router /api/users/{userid}/address [get]
func (cfg *AppConfig) GetAddressByUserid(c *fiber.Ctx) error {
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
		cfg.Logger.Error("Couldnt connect to DB: " + err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	_, _, err = db.GetUserById(userid)
	if err != nil {
		newErr.Message = "User with the given ID is not found!"
		newErr.Code = fiber.StatusNotFound
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	userAddress, err := db.GetUserAddressById(userid)
	if err != nil {
		newErr.Message = "No Address found for the user!"
		newErr.Code = fiber.StatusNotFound
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = userAddress
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetUsers method for getting all existing users.
// @Description Get all existing users.
// @Summary Get all existing users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.RequestResponse "OK"
// @Failure 403 {object} response.ErrorResponse{message=string,int} "FORBIDDEN"
// @Failure 404 {object} response.ErrorResponse{message=string,int} "NOT FOUND"
// @Failure 503 {object} response.ErrorResponse{message=string,int} "SERVICE UNAVAILABLE"
// @Security BearerAuth
// @Router /api/users [get]
func (cfg *AppConfig) GetUsers(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	//Get claims from JWT.
	data := cfg.JwtCredentials(c)
	isAdmin := data["isAdmin"].(bool)

	if !isAdmin {
		newErr.Message = "InsufficientPermissions. Admin rights needed"
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
		newErr.Code = fiber.StatusNotFound
		cfg.Logger.Error("Couldnt connect to DB: " + err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Define users variable.
	result := response.GetUsersResponse{}
	for _, user := range users {
		result.Users = append(result.Users, response.GetUserResponse{Userid: user.Userid,
			LastModifiedAt:   user.AuditInfo.LastModifiedAt,
			CreatedAt:        user.AuditInfo.CreatedAt,
			Username:         user.Username,
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
// @Success 200 {object} response.RequestResponse "OK"
// @Failure 400 {object} response.ErrorResponse{message=string,code=int} "BAD REQUEST"
// @Failure 404 {object} response.ErrorResponse{message=string,int} "NOT FOUND"
// @Failure 503 {object} response.ErrorResponse{message=string,int} "SERVICE UNAVAILABLE"
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
		cfg.Logger.Error("Couldnt connect to DB: " + err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	foundUser, _, err := db.GetUserById(userid)
	if err != nil {
		newErr.Message = "User with the given ID is not found!"
		newErr.Code = fiber.StatusNotFound
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
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
