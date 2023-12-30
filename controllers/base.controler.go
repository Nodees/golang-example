package controllers

import (
	connection "core/connections"
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

func BaseCreate[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		model := new(T)
		err := c.BodyParser(model)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		result := connection.DB.Create(model)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create record"})
		}
		return c.JSON(model)
	}
}

func BaseRetrieve[T any](preloads ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		model := new(T)
		query := connection.DB.Find(&model, id)

		if len(preloads) > 0 {
			for _, preload := range preloads {
				query = query.Preload(preload)
			}
		}

		result := query.First(model)
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
		result := connection.DB.Find(model, id).Delete(model)
		if result.Error != nil {
			return result.Error
		}
		return c.SendStatus(http.StatusOK)
	}
}
