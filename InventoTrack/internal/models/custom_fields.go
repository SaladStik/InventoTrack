package models

import "time"

// CustomField represents a custom tracking field
type CustomField struct {
	ID              int       `json:"id"`
	CompanyID       int       `json:"company_id"`
	TableName       string    `json:"table_name"`
	FieldName       string    `json:"field_name"`
	FieldType       string    `json:"field_type"`
	ValidationRules string    `json:"validation_rules"`
	CreatedAt       time.Time `json:"created_at"`
}
