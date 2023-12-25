package controllers

import (
	connection "core/connections"
	"core/models"

	"github.com/gofiber/fiber/v2"
)

func AddressList(c *fiber.Ctx) error {
	var addresses []models.Address

	connection.DB.Find(&addresses)

	return c.JSON(addresses)
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
	var address models.Address

	id := c.Params("id")
	connection.DB.Find(&address, id)

	if address.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Address not found")
	}

	return c.JSON(address)
}

func AddressUpdate(c *fiber.Ctx) error {
	var address models.Address
	id := c.Params("id")

	if err := connection.DB.First(&address, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := connection.DB.Model(&address).Updates(updateData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update record"})
	}

	return c.JSON(address)
}

func AddressDestroy(c *fiber.Ctx) error {
	var address models.Address

	id := c.Params("id")
	connection.DB.First(&address, id)

	if address.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Address not found")
	}

	connection.DB.Unscoped().Delete(&address)
	return c.JSON(address)
}
