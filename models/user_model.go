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
	Name      string    `gorm:"not null;size:255" json:"name"`
	Email     string    `gorm:"not null;uniqueIndex;size:255" json:"email"`
	Password  string    `gorm:"not null;size:255" json:"-"`
	Tel       string    `gorm:"not null;size:255" json:"tel"`
	Role      Role      `gorm:"not null;size:255default:user" json:"role"`
	Token     string    `gorm:"size:255" json:"token,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

type Token struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}
