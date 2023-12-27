package controllers

import (
	connection "core/connections"
	"core/models"

	"github.com/gofiber/fiber/v2"
)

func AddressList(c *fiber.Ctx) error {
	result := BaseList[models.Address]()
	return result(c)
}

func AddressCreate(c *fiber.Ctx) error {
	address := new(models.Address)

	if err := c.BodyParser(address); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	connection.DB.Create(&address)
	return c.JSON(address)
}

func AddressRetrieve(c *fiber.Ctx) error {
	result := BaseRetrieve[models.Address]()
	return result(c)
}

func AddressUpdate(c *fiber.Ctx) error {
	result := BaseUpdate[models.Address]()
	return result(c)
}

func AddressDestroy(c *fiber.Ctx) error {
	result := BaseDestroy[models.Address]()
	return result(c)
}
