package controllers

import (
	"core/models"

	"github.com/gofiber/fiber/v2"
)

func AddressList(c *fiber.Ctx) error {
	result := BaseList[models.Address]()
	return result(c)
}

func AddressCreate(c *fiber.Ctx) error {
	result := BaseCreate[models.Address]()
	return result(c)
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
