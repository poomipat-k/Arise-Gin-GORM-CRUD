package database

import (
	"fmt"
	"log"

	"github.com/poomipat-k/crud-arise/internal/config"
	"github.com/poomipat-k/crud-arise/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func AutoMigrate() {
	DB().AutoMigrate(&models.Item{})
}

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.PostgresHost,
		config.AppConfig.PostgresUser,
		config.AppConfig.PostgresPassword,
		config.AppConfig.PostgresDb,
		config.AppConfig.PostgresPort,
	)

	var dbErr error
	db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("Failed to connect to the database: %v", dbErr)
	}
	log.Println("Database connected successfully!")
}
