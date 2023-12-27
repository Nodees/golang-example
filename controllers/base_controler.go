package controllers

import (
	connection "core/connections"
	"core/models"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BaseList[T any](preloads ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		models := new([]T)
		query := connection.DB
		if len(preloads) > 0 {
			for _, preload := range preloads {
				query = query.Preload(preload)
			}
		}
		result := query.Find(models)
		if result.Error != nil {
			return result.Error
		}
		c.SendStatus(http.StatusOK)
		return c.JSON(models)
	}
}

func BaseCreate(c *fiber.Ctx) error {
	address := new(models.Address)

	if err := c.BodyParser(address); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	connection.DB.Create(&address)
	return c.JSON(address)
}

func BaseRetrieve[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		model := new(T)
		result := connection.DB.First(model, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
			}
			return result.Error
		}
		return c.JSON(model)
	}
}

func BaseUpdate[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		model := new(T)
		err := c.BodyParser(model)
		if err != nil {
			return err
		}
		result := connection.DB.Model(model).Where("id = ?", id).Updates(model)
		if result.Error != nil {
			return result.Error
		}
		return c.SendStatus(http.StatusOK)
	}
}

func BaseDestroy[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		model := new(T)
		result := connection.DB.Where("id = ?", id).Delete(model)
		if result.Error != nil {
			return result.Error
		}
		return c.SendStatus(http.StatusOK)
	}
}
