package routers

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	UserSetupRouter(app)
	AddressSetupRouter(app)
	AuthRouter(app)
}
