package main

import (
	"core/config"
	connection "core/connections"
	"core/models"
	"core/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	loadConfig, _ := config.LoadConfig(".")

	corsConfig := config.CorsConfig(&loadConfig)
	app.Use(corsConfig)

	connection.InitPostgresDB(&loadConfig)
	models.Migrate(connection.DB)

	cas := config.CasbinConfig(&loadConfig)
	routes.Routes(app, cas)

	log.Fatal(app.Listen(":8000"))
}
