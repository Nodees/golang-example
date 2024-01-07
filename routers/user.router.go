package routers

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title GoFiber Example API
// @version 1.0
// @description Golang GoFiber swagger auto generate step by step by swaggo
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func UserSetupRouter(app *fiber.App) {
	app.Get("/api/swagger/*", swagger.HandlerDefault) // default
	app.Get("/api/user", controllers.UserList)
	app.Post("/api/user", controllers.UserCreate)
	app.Get("/api/user/:id", controllers.UserRetrive)
	app.Delete("/api/user/:id", controllers.UserDestroy)
	app.Patch("/api/user/:id", controllers.UserUpdate)
}
