package services

import (
	"errors"
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"

	"golang.org/x/crypto/bcrypt"
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

// CreateCompanyWithOwnerService handles the creation of a company and its owner
func CreateCompanyWithOwnerService(companyName, username, email, password string) (*models.Company, *models.User, error) {
	if companyName == "" || username == "" || email == "" || password == "" {
		return nil, nil, errors.New("all fields are required")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, errors.New("failed to hash password")
	}

	// Create the company
	company := &models.Company{Name: companyName}

	// Create the owner user
	owner := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword), // Use the hashed password
		Role:         "owner",
	}

	// Create the company and owner in the database
	err = repositories.CreateCompanyWithOwner(company, owner)
	if err != nil {
		return nil, nil, err
	}

	return company, owner, nil
}
