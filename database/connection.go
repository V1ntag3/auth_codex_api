package database

import (
	"auth_codex_api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(database string) {
	// connection to database
	connection, err := gorm.Open(sqlite.Open(database), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	} else {
		log.Println("Database connected")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.App{})

	connection.AutoMigrate(&models.ValidToken{})

}
