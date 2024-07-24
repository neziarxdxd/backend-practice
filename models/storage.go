package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the he database connection.
var DB *gorm.DB

// SetupDatabase migrates and sets up the database.
func SetupDatabase() {
	// u := helper.GetEnv("DATABASE_USER", "golang")
	// p := helper.GetEnv("DATABSE_PASSWORD", "golang")
	// h := helper.GetEnv("DATABASE_HOST", "localhost:3306")
	// n := helper.GetEnv("DATABASE_NAME", "go_test")
	// q := "charset=utf8mb4&parseTime=True&loc=Local"

	// // Assemble the connection string.
	// // import "gorm.io/driver/postgres"
	// // ref: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	dsn := "user=postgres password= dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrate the schema
	tx := db

	if err := tx.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate database User: " + err.Error())
	}

	if err := tx.AutoMigrate(&Food{}); err != nil {
		panic("failed to migrate database Food: " + err.Error())
	}

	if err != nil {
		fmt.Println("Could not open database connection")
	} else {
		fmt.Println("Testing huhuhu")
	}

	DB = db
}
