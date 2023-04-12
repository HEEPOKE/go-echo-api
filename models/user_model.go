package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`
	Tel      string `gorm:"not null"`
	Role     Role   `gorm:"not null;default:user"`
}
