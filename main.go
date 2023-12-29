package main

import (
	"core/configs"
	"core/configs/middleware"
	connection "core/connections"
	"core/models"
	"core/routers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	loadConfig, _ := configs.LoadConfig(".")

	corsConfig := configs.CorsConfig(&loadConfig)
	app.Use(corsConfig)
	app.Use(middleware.Authenticate(&loadConfig))

	connection.InitPostgresDB(&loadConfig)
	models.Migrate(connection.DB)

	// cas := config.Authenticate(&loadConfig)
	routers.Init(app)

	log.Fatal(app.Listen(":8000"))
}
