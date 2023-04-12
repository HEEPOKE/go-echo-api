package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;uniqueIndex"`
	Password  string `gorm:"not null"`
	Tel       string `gorm:"not null"`
	Role      Role   `gorm:"not null;default:user"`
	Token     string
	ExpiresAt time.Time
}

type Token struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}
