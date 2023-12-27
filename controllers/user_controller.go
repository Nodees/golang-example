package controllers

import (
	"core/models"

	"github.com/gofiber/fiber/v2"
)

func UserList(c *fiber.Ctx) error {
	result := BaseList[models.User]("Address")
	return result(c)
}

func UserRetrive(c *fiber.Ctx) error {
	result := BaseRetrieve[models.User]()
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
