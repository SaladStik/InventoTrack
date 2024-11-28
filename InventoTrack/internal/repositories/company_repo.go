package repositories

import (
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"

	"gorm.io/gorm"
)

// CreateCompany saves a new company to the database
func CreateCompany(company *models.Company) error {
	result := extensions.DB.Create(company)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateCompanyWithOwner creates a new company and its owner in a single transaction
func CreateCompanyWithOwner(company *models.Company, owner *models.User) error {
	return extensions.DB.Transaction(func(tx *gorm.DB) error {
		// Step 1: Create the company
		if err := tx.Create(company).Error; err != nil {
			return err
		}

		// Step 2: Create the owner and associate with the company
		companyID := company.ID // Assign the ID to a variable
		owner.CompanyID = &companyID
		owner.Role = "owner"
		if err := tx.Create(owner).Error; err != nil {
			return err
		}

		// Step 3: Update the company with the owner's ID
		ownerID := owner.ID // Use a local variable to take the address
		if err := tx.Model(&models.Company{}).Where("id = ?", company.ID).Update("owner_id", &ownerID).Error; err != nil {
			return err
		}

		return nil
	})
}
