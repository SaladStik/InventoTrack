package services

import (
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"
)

// GetRecycleBinEntries fetches all recycle bin entries for a company
func GetRecycleBinEntries(companyID int) ([]models.RecycleBin, error) {
	return repositories.GetRecycleBinEntries(companyID)
}

// PermanentlyDeleteEntry permanently deletes an entry from the recycle bin
func PermanentlyDeleteEntry(id int) error {
	// Check if the entry exists
	entry, err := repositories.GetRecycleBinEntryByID(id)
	if err != nil {
		return err
	}

	// Permanently delete the record
	err = repositories.PermanentlyDeleteRecord(entry.TableName, entry.RecordID)
	if err != nil {
		return err
	}

	// Remove the entry from the recycle bin
	return repositories.DeleteRecycleBinEntry(id)
}

// AddToRecycleBin adds a record to the recycle bin (for internal use)
func AddToRecycleBin(companyID int, tableName string, recordID int) error {
	entry := models.RecycleBin{
		CompanyID:       companyID,
		TableName:       tableName,
		RecordID:        recordID,
		RetentionPeriod: 30, // Default retention period
	}
	return repositories.AddRecycleBinEntry(entry)
}
