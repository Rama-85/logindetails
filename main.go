package main

import (
	"logindetails/controller"
	"logindetails/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	models.ConnectDB() // new
	r.POST("/login", controller.CreateLogin)
	r.Run("localhost:8081")
	err := r.Run()
	if err != nil {
		return
	}

}
