package main

import (
	"logindetails/config"
	"logindetails/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB() // new
	r.POST("/login", controller.CreateLogin)
	r.Run()
}
