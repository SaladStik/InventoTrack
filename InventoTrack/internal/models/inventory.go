package models

import "time"

// Inventory represents an inventory item
type Inventory struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	CompanyID  int       `json:"company_id"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
