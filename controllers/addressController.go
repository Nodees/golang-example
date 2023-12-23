package controllers

import (
	"net/http"

	connection "core/connections"
	"core/models"

	"github.com/gin-gonic/gin"
)

func AddressList(c *gin.Context) {
	var address []models.Address

	connection.DB.Find(&address)

	c.JSON(http.StatusOK, gin.H{
		"results": address,
	})
}
