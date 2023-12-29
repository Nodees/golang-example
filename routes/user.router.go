package routes

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserSetupRoutes(app *fiber.App) {
	app.Get("/api/user", controllers.UserList)
	app.Post("/api/user/login", controllers.LoginHandler)
	app.Get("/api/user/logout", controllers.Logout)
	app.Post("/api/user", controllers.UserCreate)
	app.Get("/api/user/:id", controllers.UserRetrive)
	app.Delete("/api/user/:id", controllers.UserDestroy)
	app.Patch("/api/user/:id", controllers.UserUpdate)
}
