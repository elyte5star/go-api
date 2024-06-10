package service

import (
	"fmt"

	"github.com/api/common/config"
	"github.com/api/common/database"
	"github.com/api/repository/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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
