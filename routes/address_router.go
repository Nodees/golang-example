package routes

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func AddressSetupRoutes(app *fiber.App) {
	app.Get("/api/address", controllers.AddressList)
	app.Post("/api/address", controllers.AddressCreate)
	app.Get("/api/address/:id", controllers.AddressRetrieve)
	app.Delete("/api/address/:id", controllers.AddressDestroy)
}
