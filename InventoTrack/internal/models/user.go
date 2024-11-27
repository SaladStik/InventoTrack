package models

import "time"

// User represents a user in the system
type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Store hashed passwords
	CompanyID    int       `json:"company_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
