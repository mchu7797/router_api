package models

import (
	"gorm.io/gorm"
)

type NetworkConfig struct {
	gorm.Model
	Id          int64
	ConfigName  string
	ConfigValue string
}
