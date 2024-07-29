package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int64
	LoginId   string
	Password  string
	Name      string
	Privilage int32
}
