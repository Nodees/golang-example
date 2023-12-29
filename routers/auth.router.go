package routers

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App) {
	app.Post("/api/login", controllers.LoginHandler)
	app.Get("/api/logout", controllers.Logout)
}
