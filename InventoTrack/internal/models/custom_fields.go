package models

import "time"

// CustomField represents a custom field definition
type CustomField struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	CompanyID       int       `json:"company_id" gorm:"not null"`
	TableName       string    `json:"table_name" gorm:"not null"`
	FieldName       string    `json:"field_name" gorm:"not null"`
	FieldType       string    `json:"field_type" gorm:"not null"` // text, number, date, boolean
	ValidationRules string    `json:"validation_rules"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// CustomFieldValue represents a value assigned to a custom field for a specific record
type CustomFieldValue struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	CustomFieldID int       `json:"custom_field_id"` // Foreign key to CustomFields table
	InventoryID   int       `json:"inventory_id"`    // Foreign key to Inventory table
	Value         string    `json:"value"`           // The value for the custom field
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
