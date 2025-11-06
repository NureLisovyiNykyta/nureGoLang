package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "school-journal/docs" // Swagger docs

	"school-journal/database"
	"school-journal/models"
	"school-journal/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title School Journal API
// @version 1.0
// @description REST API for managing school journal system
// @termsOfService http://swagger.io/terms/

// @contact.name Nikita Lisovyi
// @contact.url https://github.com/yourusername
// @contact.email your.email@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: .env file not found, using system environment variables")
	}

	// Connect to DB
	if err := database.Connect(); err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	// Auto migrate models
	if err := database.DB.AutoMigrate(
		&models.Subject{},
		&models.Teacher{},
		&models.Class{},
		&models.Student{},
		&models.Lesson{},
		&models.Grade{},
	); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}
	log.Println("✅ AutoMigrate completed successfully")

	// Initialize Gin router
	r := gin.Default()

	// Register all routes
	routes.InitRoutes(r)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	log.Println("🚀 Server is running at http://localhost:8080")
	log.Println("📘 Swagger docs available at http://localhost:8080/swagger/index.html")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
