package database

import (
	"github.com/mchu7797/router_api/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect(databasePath string) error {
	var err error

	Database, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})

	if err != nil {
		return err
	}

	Database.AutoMigrate(&models.IntranetConfig{})
	Database.AutoMigrate(&models.NetworkConfig{})
	Database.AutoMigrate(&models.User{})

	return nil
}
