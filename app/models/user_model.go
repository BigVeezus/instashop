package models

import "gorm.io/gorm"

// User model definition
type User struct {
	gorm.Model
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	IsAdmin   bool   `json:"is_admin" gorm:"default:false"`
	Email     string `json:"email" gorm:"unique" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"is_admin" gorm:"default:false"`
}
