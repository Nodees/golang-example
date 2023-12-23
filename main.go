package main

import (
	connection "core/connections"
	"core/models"
	"core/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	connection.InitDB()
	connection.DB.AutoMigrate(&models.User{}, &models.Address{})

	routes.SetupRoute(app)

	log.Fatal(app.Listen(":8000"))
}
