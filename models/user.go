package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id             int    `json:"ID" gorm:"primary_key"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Login          string `json:"login"`
	ForgotPassword string `json:"forgot password"`
}

type CreateLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
