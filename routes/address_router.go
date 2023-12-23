package routes

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func AddressSetupRoutes(app *fiber.App) {
	app.Get("/api/address", controllers.AddressList)
}
