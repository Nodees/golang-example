package routers

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserSetupRouter(app *fiber.App) {
	app.Get("/api/user", controllers.UserList)
	app.Post("/api/user", controllers.UserCreate)
	app.Get("/api/user/:id", controllers.UserRetrive)
	app.Delete("/api/user/:id", controllers.UserDestroy)
	app.Patch("/api/user/:id", controllers.UserUpdate)
}
