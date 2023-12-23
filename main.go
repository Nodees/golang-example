package main

import (
	connection "core/connections"
	"core/controllers"
	"core/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	connection.InitDB()
	connection.DB.AutoMigrate(&models.User{}, &models.Address{})

	r.GET("/users", controllers.UserList)
	r.GET("/address", controllers.AddressList)

	r.Run()
}
