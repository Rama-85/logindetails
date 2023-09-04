package controller

import (
	"net/http"

	"github.com/Rama-85/logindetails/models"
	"github.com/gin-gonic/gin"
)

func CreateLogin(c *gin.Context) {
	// Validate input
	var input CreateLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create login
	login := models.User{Email: input.Email, Password: input.Password}
	models.DB.Create(&login)

	c.JSON(http.StatusOK, gin.H{"data": login})
}
