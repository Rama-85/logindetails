package controller

import (
	"net/http"

	"logindetails/models"

	"github.com/gin-gonic/gin"
)

type CreateLoginInput struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	ForgotPassword string `json:"forgot password" binding:"required"`
}

func CreateLogin(c *gin.Context) {

	// Validate input
	var input CreateLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create login
	login := models.User{Name: input.Name, Email: input.Email, Password: input.Password, ForgotPassword: input.ForgotPassword}
	models.DB.Create(&login)

	c.JSON(http.StatusOK, gin.H{"data": login})
}
