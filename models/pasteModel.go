package models

import "gorm.io/gorm"

type Paste struct {
	gorm.Model
	ID   uint
	User User `gorm:"foreignKey:ID"`
}

type User struct {
	gorm.Model
	ID  uint
	Ips []string
}
