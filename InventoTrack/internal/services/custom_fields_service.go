package services

import (
	"errors"
	"inventotrack/internal/repositories"
)

// AddCustomField adds a new custom field for a company
func AddCustomField(tableName, fieldName, fieldType, validationRules string, companyID int) error {
	// Validate input
	if tableName == "" || fieldName == "" || fieldType == "" {
		return errors.New("table name, field name, and field type are required")
	}

	// Call the repository to add the custom field
	err := repositories.AddCustomField(tableName, fieldName, fieldType, validationRules, companyID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveCustomField removes a custom field by moving it to the recycle bin
func RemoveCustomField(id int) error {
	// Call the repository to remove the custom field
	err := repositories.RemoveCustomField(id)
	if err != nil {
		return err
	}

	return nil
}
