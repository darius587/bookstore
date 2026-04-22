package db

import (
	"bookstore2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=201105 dbname=bookstore port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	DB.AutoMigrate(
		&models.User{},
		&models.Author{},
		&models.Book{},
		&models.Category{},
		&models.Favorite{},
	)

	log.Println("DB connected")
}
