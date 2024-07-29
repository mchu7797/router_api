package database

import (
	"github.com/mchu7797/router_api/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(databasePath string) {
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database!")
	}

	db.AutoMigrate(&models.IntranetConfig{})
	db.AutoMigrate(&models.NetworkConfig{})
	db.AutoMigrate(&models.User{})
}
