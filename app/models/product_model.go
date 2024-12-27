package models

import "gorm.io/gorm"

// Product model
type Product struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}
