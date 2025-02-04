package config

import (
	"awesomeProject/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=user password=1234 dbname=go_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Nu s-a putut conecta la baza de date:", err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.AuthRequestUser{})
	DB = db
	fmt.Println("✅ Conectat la baza de date!")
}
