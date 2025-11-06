package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect opens a connection to Postgres (Azure) using env vars and configures GORM.
func Connect() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	schema := os.Getenv("DB_SCHEMA")
	
	if sslmode == "" {
		sslmode = "require"
	}

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return fmt.Errorf("database connection env variables are not fully set")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=%s",
		host, user, password, dbname, port, sslmode, schema,
	)

	// Use a moderate logger for migrations and queries
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // ignore ErrRecordNotFound errors
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	log.Println("✅ Connected to PostgreSQL (Azure)")

	// Set schema search path
	if err := DB.Exec(fmt.Sprintf("SET search_path TO %s", schema)).Error; err != nil {
		log.Fatalf("❌ Failed to set schema: %v", err)
	}

	log.Printf("📘 Using schema: %s", schema)

	// Create DB extensions needed for uuid generation if not exists.
	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`).Error; err != nil {
		// not fatal, but log
		log.Printf("⚠️  Warning: could not ensure pgcrypto extension: %v\n", err)
	} else {
		log.Println("✅ Ensured pgcrypto extension exists")
	}
	
	return nil
}
