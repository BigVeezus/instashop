package config

import (
	"fmt"
	"go-gin-app/app/models"

	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Initialize database connection

func SetupDatabase() {

	// Get the database URI from environment variables
	dbURI := os.Getenv("DB_URI")

	if dbURI == "" {
		log.Fatal("DB_URI environment variable not set")
	}

	// Open the database connection using GORM v2
	var err error
	DB, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	// Handle any errors that occur while connecting to the database
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Optionally, you can check if the database is reachable
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error getting SQL DB instance:", err)
	}

	// Set the max idle connections and max open connections for better connection management
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Run migrations - this will automatically create tables and update the schema
	err = DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}

	fmt.Println("Database connected and migrated!")
}
