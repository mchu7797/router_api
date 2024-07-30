package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int64  `json:"id"`
	LoginId   string `json:"login_id"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Privilage int32  `json:"privilage"`
}
