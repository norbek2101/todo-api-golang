package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "todo-api-golang/models"
)

var DB *gorm.DB

func Connect() {
    // Replace the connection details with your PostgreSQL information
	dbURL := "host=localhost port=5432 user=user_service dbname=user_service password=Qwerty123$ sslmode=disable"
    db, err := gorm.Open("postgres", dbURL)
    if err != nil {
        panic("Failed to connect to the database")
    }

    DB = db
    DB.AutoMigrate(&models.Todo{})
}
