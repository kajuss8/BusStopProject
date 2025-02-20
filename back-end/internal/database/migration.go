package database

import (
	"busProject/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Stop{}, &models.Route{}, &models.Trip{}, &models.StopTime{}, &models.Calendar{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	log.Println("Database migrated successfully.")
	return nil
}


