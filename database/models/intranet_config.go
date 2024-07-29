package models

import (
	"gorm.io/gorm"
)

type IntranetConfig struct {
	gorm.Model
	Id          int64
	ConfigName  string
	ConfigValue string
}
