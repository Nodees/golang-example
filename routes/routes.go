package routes

import (
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, authz *casbin.Middleware) {
	UserSetupRoutes(app, authz)
	AddressSetupRoutes(app, authz)
}
