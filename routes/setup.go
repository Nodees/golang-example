package routes

import (
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App, authz *casbin.Middleware) {
	UserSetupRoutes(app, authz)
	AddressSetupRoutes(app, authz)
}
