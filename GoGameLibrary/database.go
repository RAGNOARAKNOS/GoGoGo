package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	databaseConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		CFG.DBHost,
		CFG.DBUser,
		CFG.DBPassword,
		CFG.DBName,
		CFG.DBPort)

	db, err := gorm.Open(postgres.Open(databaseConnection), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	// Use automigrate to create the 'games' table - if one doesn't already exist
	DB.AutoMigrate(&Game{})
}
