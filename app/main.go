package main

import (
	"go-gin-app/app/config"
	"go-gin-app/app/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize database
	config.SetupDatabase()

	// Create a Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
