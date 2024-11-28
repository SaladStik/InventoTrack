package models

import "time"

// Company represents a company in the system
type Company struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	OwnerID   *int   `gorm:"unique"` // Pointer to allow null values
	CreatedAt time.Time
	UpdatedAt time.Time
}
