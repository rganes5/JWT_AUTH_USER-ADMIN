package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"index:idx_name,unique"`
	Password string
	Number   string `gorm:"unique"`
	Role     string `gorm:"not null;default:user"`
}
