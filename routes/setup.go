package routes

import "github.com/gofiber/fiber/v2"

func SetupRoute(app *fiber.App) {
	UserSetupRoutes(app)
	AddressSetupRoutes(app)
}
