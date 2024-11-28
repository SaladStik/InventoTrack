package models

import "time"

type Log struct {
	ID        int    `gorm:"primaryKey"`
	Action    string `gorm:"not null"`
	UserID    int    `gorm:"not null"`
	Entity    string `gorm:"not null"`
	EntityID  int    `gorm:"not null"`
	Details   string `gorm:"type:text"`
	CreatedAt time.Time
}
