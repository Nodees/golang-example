package routers

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func AddressSetupRouter(app *fiber.App) {
	app.Get("/api/address", controllers.AddressList)
	app.Post("/api/address", controllers.AddressCreate)
	app.Get("/api/address/:id", controllers.AddressRetrieve)
	app.Delete("/api/address/:id", controllers.AddressDestroy)
	app.Patch("/api/address/:id", controllers.AddressUpdate)
}
