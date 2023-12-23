package routes

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserSetupRoutes(app *fiber.App) {
	app.Get("/api/users", controllers.UserList)
}
