package repositories

import (
	"errors"
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
)

// AddCustomField adds a custom field to the database
func AddCustomField(field *models.CustomField) error {
	result := extensions.DB.Create(field)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// RemoveCustomField removes a custom field and moves it to the recycle bin
func RemoveCustomField(id int) error {
	var customField models.CustomField

	// Fetch the custom field
	result := extensions.DB.First(&customField, id)
	if result.Error != nil {
		return errors.New("custom field not found")
	}

	// Move the custom field to the recycle bin
	recycleBinEntry := models.RecycleBin{
		CompanyID:       customField.CompanyID,
		TableName:       "Custom_Fields",
		RecordID:        id,
		DeletedAt:       customField.CreatedAt,
		RetentionPeriod: 30,
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

// GetCustomFields fetches all custom fields for a company and a table
func GetCustomFields(companyID int, tableName string) ([]models.CustomField, error) {
	var customFields []models.CustomField
	err := extensions.DB.Where("company_id = ? AND table_name = ?", companyID, tableName).Find(&customFields).Error
	return customFields, err
}

// AddCustomFieldValue adds a custom field value to the database
func AddCustomFieldValue(customFieldID int, inventoryID int, value string) error {
	customFieldValue := models.CustomFieldValue{
		CustomFieldID: customFieldID,
		InventoryID:   inventoryID,
		Value:         value,
	}

	result := extensions.DB.Create(&customFieldValue)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetCustomFieldValues fetches all custom field values for a specific record
func GetCustomFieldValues(entityID int) ([]models.CustomFieldValue, error) {
	var values []models.CustomFieldValue
	err := extensions.DB.Where("entity_id = ?", entityID).Find(&values).Error
	return values, err
}
