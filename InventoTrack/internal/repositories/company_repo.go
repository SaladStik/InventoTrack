package repositories

import (
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
)

// CreateCompany saves a new company to the database
func CreateCompany(company *models.Company) error {
	result := extensions.DB.Create(company)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
