package repositories

import (
	"errors"
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
)

// AddCustomField adds a custom field to the database
func AddCustomField(tableName, fieldName, fieldType, validationRules string, companyID int) error {
	// Insert the custom field into the CustomFields table
	customField := models.CustomField{
		CompanyID:       companyID,
		TableName:       tableName,
		FieldName:       fieldName,
		FieldType:       fieldType,
		ValidationRules: validationRules,
	}

	result := extensions.DB.Create(&customField)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// RemoveCustomField removes a custom field and moves it to the recycle bin
func RemoveCustomField(id int) error {
	// Fetch the custom field
	var customField models.CustomField
	result := extensions.DB.First(&customField, id)
	if result.Error != nil {
		return errors.New("custom field not found")
	}

	// Move the custom field to the recycle bin
	recycleBinEntry := models.RecycleBin{
		CompanyID:       customField.CompanyID,
		TableName:       "CustomFields",
		RecordID:        id,
		DeletedAt:       customField.CreatedAt, // Track the deletion time
		RetentionPeriod: 30,                    // Default retention period
	}

	if err := extensions.DB.Create(&recycleBinEntry).Error; err != nil {
		return err
	}

	// Delete the custom field
	if err := extensions.DB.Delete(&customField).Error; err != nil {
		return err
	}

	return nil
}
