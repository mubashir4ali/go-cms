package database

import (
	"cms-api/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// dsn := "host=localhost user=postgres password=postgres dbname=cms port=5432 sslmode=disable"
	// dsn := "host=db user=postgres password=postgres dbname=cms port=5432 sslmode=disable"
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable search_path=cms"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
	DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Content{})
	fmt.Println("Database connected and migrated successfully")
}
