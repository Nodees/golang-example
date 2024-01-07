package controllers

import (
	"core/models"

	"github.com/gofiber/fiber/v2"
)

func UserList(c *fiber.Ctx) error {
	result := BaseList[models.User]("Address")
	return result(c)
}

// HandlerCreateUser godoc
//
//	@Summary		Create new user
//	@Description	Create new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		models.User	true	"User"
//	@Success		200		{string}  string  "OK"
//	@Failure		400		{string}  error  "Bad Request"
//	@Router			/api/user [post]
func UserCreate(c *fiber.Ctx) error {
	result := BaseCreate[models.User]()
	return result(c)
}

func UserRetrive(c *fiber.Ctx) error {
	result := BaseRetrieve[models.User]("Address")
	return result(c)
}

func UserUpdate(c *fiber.Ctx) error {
	result := BaseUpdate[models.User]()
	return result(c)
}

func UserDestroy(c *fiber.Ctx) error {
	result := BaseDestroy[models.User]()
	return result(c)
}
