package controllers

import (
	connection "core/connections"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PaginationParams struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

func BaseList[T any](preloads ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		models := new([]T)
		query := connection.DB
		params := new(PaginationParams)

		if err := c.QueryParser(params); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
		}

		if params.Page != 0 || params.PageSize != 0 {
			return BasePaginatedList[T](preloads...)(c)
		}

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

func BasePaginatedList[T any](preloads ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		models := new([]T)
		query := connection.DB
		params := new(PaginationParams)
		var totalCount int64

		if err := c.QueryParser(params); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
		}

		if params.Page <= 0 {
			params.Page = 1
		}

		defaultPageSize := 10
		if params.PageSize <= 0 {
			params.PageSize = defaultPageSize
		}

		if len(preloads) > 0 {
			for _, preload := range preloads {
				query = query.Preload(preload)
			}
		}

		if err := query.Model(models).Count(&totalCount).Error; err != nil {
			return err
		}

		totalPages := int(totalCount / int64(params.PageSize))
		if totalCount%int64(params.PageSize) > 0 {
			totalPages++
		}

		if params.Page == totalPages && totalCount%int64(params.PageSize) == 0 {
			totalPages--
		}

		offset := (params.Page - 1) * params.PageSize
		query = query.Offset(offset).Limit(params.PageSize)

		result := query.Find(models)
		if result.Error != nil {
			return result.Error
		}

		response := fiber.Map{
			"total_pages":  totalPages,
			"current_page": params.Page,
			"page_size":    params.PageSize,
			"result":       models,
		}

		c.Status(fiber.StatusOK)
		return c.JSON(response)
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
