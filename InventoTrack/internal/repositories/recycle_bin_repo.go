package repositories

import (
	"errors"
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
)

// GetRecycleBinEntries retrieves all entries in the recycle bin for a company
func GetRecycleBinEntries(companyID int) ([]models.RecycleBin, error) {
	var entries []models.RecycleBin
	result := extensions.DB.Where("company_id = ?", companyID).Find(&entries)
	if result.Error != nil {
		return nil, result.Error
	}
	return entries, nil
}

// GetRecycleBinEntryByID retrieves a single recycle bin entry by its ID
func GetRecycleBinEntryByID(id int) (models.RecycleBin, error) {
	var entry models.RecycleBin
	result := extensions.DB.First(&entry, id)
	if result.Error != nil {
		return entry, errors.New("recycle bin entry not found")
	}
	return entry, nil
}

// AddRecycleBinEntry adds a record to the recycle bin
func AddRecycleBinEntry(entry models.RecycleBin) error {
	result := extensions.DB.Create(&entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// PermanentlyDeleteRecord permanently deletes a record from a specific table
func PermanentlyDeleteRecord(tableName string, recordID int) error {
	query := "DELETE FROM " + tableName + " WHERE id = ?"
	result := extensions.DB.Exec(query, recordID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteRecycleBinEntry deletes an entry from the recycle bin
func DeleteRecycleBinEntry(id int) error {
	result := extensions.DB.Delete(&models.RecycleBin{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
