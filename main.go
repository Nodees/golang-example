package main

import (
	postgres "core/connections"
	models "core/models"
)

func main() {
	postgres.InitDB()
	postgres.DB.AutoMigrate(&models.User{})
}
