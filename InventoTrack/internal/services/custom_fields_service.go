package services

import (
	"errors"
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"
)

// AddCustomField adds a new custom field for a company
func AddCustomField(companyID int, tableName, fieldName, fieldType, validationRules string) error {
	// Validate input
	if tableName == "" || fieldName == "" || fieldType == "" {
		return errors.New("table name, field name, and field type are required")
	}

	// Create the custom field model
	field := &models.CustomField{
		CompanyID:       companyID,
		TableName:       tableName,
		FieldName:       fieldName,
		FieldType:       fieldType,
		ValidationRules: validationRules,
	}

	return repositories.AddCustomField(field)
}

// RemoveCustomField removes a custom field and moves it to the recycle bin
func RemoveCustomField(id int) error {
	return repositories.RemoveCustomField(id)
}

// GetCustomFields fetches all custom fields for a company and table
func GetCustomFields(companyID int, tableName string) ([]models.CustomField, error) {
	return repositories.GetCustomFields(companyID, tableName)
}

// AddCustomFieldValue assigns a value to a custom field for a specific entity
func AddCustomFieldValue(customFieldID int, inventoryID int, value string) error {
	if customFieldID <= 0 || inventoryID <= 0 || value == "" {
		return errors.New("custom field ID, inventory ID, and value are required")
	}

	return repositories.AddCustomFieldValue(customFieldID, inventoryID, value)
}

// GetCustomFieldValues fetches all custom field values for an entity
func GetCustomFieldValues(inventoryID int) ([]models.CustomFieldValue, error) {
	var values []models.CustomFieldValue
	result := extensions.DB.Where("inventory_id = ?", inventoryID).Find(&values)
	return values, result.Error
}
