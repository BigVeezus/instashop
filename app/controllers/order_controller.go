package controllers

import (
	"go-gin-app/app/config"
	"go-gin-app/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Place an order
func PlaceOrder(c *gin.Context) {
	var order models.Order
	var input struct {
		UserID     uint   `json:"user_id"`
		Status     string `json:"status"`
		OrderItems []struct {
			ProductID uint `json:"product_id" binding:"required"`
			Quantity  uint `json:"quantity" binding:"required"`
		} `json:"order_items" binding:"required"`
	}

	// Bind and validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(float64)

	// Begin transaction
	tx := config.DB.Begin()

	// Initialize the order
	order.UserID = uint(userID)
	order.Status = "Pending"

	// Save the order
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Initialize total cost
	totalCost := 0.0

	// Process each order item
	for _, item := range input.OrderItems {
		var product models.Product
		if err := config.DB.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		// Calculate subtotal for the order item
		subTotal := float64(item.Quantity) * product.Price

		// Create an order item
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			SubTotal:  subTotal,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order item"})
			return
		}

		// Add to total cost
		totalCost += subTotal
	}

	// Update order's total cost
	order.TotalCost = totalCost
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order total"})
		return
	}

	// Commit transaction
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order": order})
}

// List all orders for a specific user
func GetUserOrders(c *gin.Context) {
	userID := c.MustGet("user_id").(float64)
	var orders []models.Order

	if err := config.DB.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// Cancel an order
func CancelOrder(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order

	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if order.Status != "Pending" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot cancel order in this status"})
		return
	}

	order.Status = "Canceled"
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}

// Update order status (admin only)
func UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order
	status := c.DefaultQuery("status", "Pending")

	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.Status = status
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}
