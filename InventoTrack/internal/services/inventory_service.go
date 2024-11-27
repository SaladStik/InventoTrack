package services

import (
	"errors"
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"
)

// AddInventory creates a new inventory item
func AddInventory(name string, companyID, categoryID int) (*models.Inventory, error) {
	// Validate input
	if name == "" {
		return nil, errors.New("inventory name is required")
	}
	if companyID <= 0 {
		return nil, errors.New("invalid company ID")
	}

	// Create the inventory object
	inventory := &models.Inventory{
		Name:       name,
		CompanyID:  companyID,
		CategoryID: categoryID,
	}

	// Save to the database
	err := repositories.CreateInventory(inventory)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

// GetInventory fetches an inventory item by its ID
func GetInventory(id int) (*models.Inventory, error) {
	if id <= 0 {
		return nil, errors.New("invalid inventory ID")
	}

	inventory, err := repositories.GetInventoryByID(id)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

// GetInventoryByID fetches an inventory item by its ID
func GetInventoryByID(id int) (*models.Inventory, error) {
	if id <= 0 {
		return nil, errors.New("invalid inventory ID")
	}

	inventory, err := repositories.GetInventoryByID(id)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}
