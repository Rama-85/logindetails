package config

import (
	"logindetails/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/weather"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	DB = db
}
