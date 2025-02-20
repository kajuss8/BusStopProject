package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return fmt.Errorf("Error opening database: %w", err)
	}

	if DB == nil {
		log.Fatal("Database connection is nil")
		return fmt.Errorf("Database connection is nil")
	}

	err = MigrateDB(DB)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
		return fmt.Errorf("Error migrating database: %w", err)
	}

	log.Println("Database connection established successfully!")
	return nil
}