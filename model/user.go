package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:64;not null"`
	Avatar   string `gorm:"size:256"`
	Phone    string `gorm:"size:11"`
	Email    string `gorm:"size:128"`
	Password string `gorm:"size:128;not null"`
}
