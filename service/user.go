package service

import "github.com/gofiber/fiber/v2"


func GetUser(c *fiber.Ctx) {

	// userId, err := uuid.FromString(c.Param("userId"))
	// if err != nil {
	// 	helpers.RespondWithError(c, http.StatusBadRequest, "invalid user id")
	// 	return
	// }

	// user, err := repository.GetUserById(context.Background(), &userId)
	// if err != nil {
	// 	config.Logger.Errorf("Error while searching: %s", err)
	// 	helpers.RespondWithError(c, http.StatusInternalServerError, "internal server error")
	// 	return
	// }
	// if user.ID.IsNil() {
	// 	helpers.RespondWithError(c, http.StatusNotFound, "user not found")
	// 	return
	// }

	// err = utils.WriteJSON(c.Writer, c)
	// if err != nil {
	// 	helpers.RespondWithError(c, http.StatusInternalServerError, "internal server error")
	// 	return
	// }

	// c.JSON(http.StatusOK, user)
}
