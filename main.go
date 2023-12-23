package main

import (
	"core/config"
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

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Não foi possivel carregar variaveis de ambiente: ", err)
	}

	connection.InitDB(&loadConfig)
	connection.DB.AutoMigrate(&models.User{}, &models.Address{})

	routes.SetupRoute(app)

	log.Fatal(app.Listen(":8000"))
}
