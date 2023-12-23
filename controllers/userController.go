package controllers

import (
	"net/http"

	connection "core/connections"
	"core/models"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var users []models.User

	connection.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"results": users,
	})
}
