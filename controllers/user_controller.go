package controllers

import (
	connection "core/connections"

	"core/models"

	"github.com/gofiber/fiber/v2"
)

func UserList(c *fiber.Ctx) error {
	var users []models.User
	connection.DB.Find(&users)
	connection.DB.Preload("Address").Find(&users)
	return c.JSON(users)
}
