package models

import "time"

// User represents a user in the system
type User struct {
	ID           int    `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"not null"` // owner, admin, user, viewer
	CompanyID    *int   `gorm:"index"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
