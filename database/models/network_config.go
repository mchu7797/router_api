package models

import (
	"gorm.io/gorm"
)

type NetworkConfig struct {
	gorm.Model
	Id          int64  `json:"id"`
	ConfigName  string `json:"config_name"`
	ConfigValue string `json:"config_value"`
}
