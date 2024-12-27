package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id" binding:"required"`   // Foreign key to the order
	ProductID uint    `json:"product_id" binding:"required"` // Foreign key to the product
	Quantity  uint    `json:"quantity" binding:"required"`   // Quantity of the product in the order
	SubTotal  float64 `json:"subtotal"`                      // Quantity * Product Price
}

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`                                                    // Reference to the user placing the order
	Status     string      `json:"status" binding:"required,oneof=Pending Completed Canceled"` // Order status enum
	TotalCost  float64     `json:"total_cost"`                                                 // Total cost of the order
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`                      // Association with order items
}
