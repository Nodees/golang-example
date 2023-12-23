package controllers

import (
	connection "core/connections"

	"core/models"

	"github.com/gofiber/fiber/v2"
)

func AddressList(c *fiber.Ctx) error {
	var address []models.Address
	connection.DB.Find(&address)
	return c.JSON(address)
}
