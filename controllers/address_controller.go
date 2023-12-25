package controllers

import (
	connection "core/connections"
	"core/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AddressList(c *fiber.Ctx) error {
	var addresses []models.Address

	connection.DB.Find(&addresses)

	return c.JSON(addresses)
}

func AddressCreate(c *fiber.Ctx) error {
	address := new(models.Address)

	log.Fatal(fmt.Sprint(c.BodyParser(address)))

	if err := c.BodyParser(address); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	connection.DB.Create(&address)
	return c.JSON(address)
}

func AddressRetrieve(c *fiber.Ctx) error {
	var address models.Address

	id := c.Params("id")
	connection.DB.Find(&address, id)

	if address.ID == 0 {
		return c.Status(404).SendString("Address not found")
	}

	return c.JSON(address)
}

func AddressDestroy(c *fiber.Ctx) error {
	var address models.Address

	id := c.Params("id")
	connection.DB.First(&address, id)

	if address.ID == 0 {
		return c.Status(404).SendString("Address not found")
	}

	connection.DB.Unscoped().Delete(&address)
	return c.JSON(address)
}
