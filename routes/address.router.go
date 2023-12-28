package routes

import (
	"core/config/middleware"
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func AddressSetupRoutes(app *fiber.App) {
	app.Get("/api/address", middleware.DeserializeUser, controllers.AddressList)
	app.Post("/api/address", middleware.DeserializeUser, controllers.AddressCreate)
	app.Get("/api/address/:id", middleware.DeserializeUser, controllers.AddressRetrieve)
	app.Delete("/api/address/:id", middleware.DeserializeUser, controllers.AddressDestroy)
	app.Patch("/api/address/:id", middleware.DeserializeUser, controllers.AddressUpdate)
}
