// models/database.go
package models

import (
	"backend-practice/helper"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database holds the database connection
type Database struct {
	DB *gorm.DB
}

// SetupDatabase initializes and migrates the database.
func SetupDatabase() (*Database, error) {
	username := helper.GetEnv("DATABASE_USERNAME", "")
	password := helper.GetEnv("DATABASE_PASSWORD", "")
	port := helper.GetEnv("DATABASE_PORT", "")
	databaseName := helper.GetEnv("DATABASE_NAME", "")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		username, password, databaseName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not open database connection: %w", err)
	}

	if err := db.AutoMigrate(&User{}, &Food{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Successfully connected to database")
	return &Database{DB: db}, nil
}
