package routes

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserSetupRoutes(app *fiber.App) {
	app.Get("/api/users", controllers.UserList)
	app.Get("/api/user/:id", controllers.UserRetrive)
	app.Delete("/api/users/:id", controllers.UserDestroy)
	app.Patch("/api/user/:id", controllers.UserUpdate)
}
