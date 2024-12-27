package routes

import (
	"go-gin-app/app/controllers"
	"go-gin-app/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// User routes
		authorized.GET("/orders", controllers.GetUserOrders)

		// Order routes
		authorized.POST("/orders", controllers.PlaceOrder)
		authorized.PUT("/orders/:id/cancel", controllers.CancelOrder)

		// Admin routes
		admin := authorized.Group("/")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.PUT("/orders/:id/status", controllers.UpdateOrderStatus)
			admin.GET("/users", controllers.GetAllUsers)

			// Product routes
			admin.POST("/products", controllers.CreateProduct)
			admin.GET("/products/:id", controllers.GetOneProduct)
			admin.PUT("/products/:id", controllers.UpdateProduct)
			admin.DELETE("/products/:id", controllers.DeleteProduct)
		}

		authorized.GET("/products", controllers.GetProducts)

	}
}
