package database

import (
	"fmt"
	"github.com/capstone-kelompok-7/backend-disappear/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPGSDatabase(config config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Check the connection
	err = db.Raw("SELECT 1").Error
	if err != nil {
		panic("failed to ping database")
	}

	fmt.Println("Connected to PostgreSQL!")
	return db
}
