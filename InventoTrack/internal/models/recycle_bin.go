package models

import "time"

// RecycleBin represents a deleted record in the recycle bin
type RecycleBin struct {
	ID              int       `json:"id"`
	CompanyID       int       `json:"company_id"`
	TableName       string    `json:"table_name"`
	RecordID        int       `json:"record_id"`
	DeletedAt       time.Time `json:"deleted_at"`
	RetentionPeriod int       `json:"retention_period"`
}
