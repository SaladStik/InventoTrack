package services

import (
	"errors"
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"
)

// CreateCompany creates a new company in the database
func CreateCompany(name string) (*models.Company, error) {
	// Validate input
	if name == "" {
		return nil, errors.New("company name is required")
	}

	// Create company object
	company := &models.Company{Name: name}

	// Save to the database
	err := repositories.CreateCompany(company)
	if err != nil {
		return nil, err
	}

	return company, nil
}
