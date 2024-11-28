package services

import (
	"context"
	"errors"
	"inventotrack/internal/middleware"
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"
)

// AddInventory creates a new inventory item
func AddInventory(ctx context.Context, name string, companyID int, categoryID *int) (*models.Inventory, error) {
	// Extract userID from context
	userID, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve user ID from context")
	}

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
		IsArchived: false, // Default to not archived
	}

	// Save to the database
	err = repositories.CreateInventory(inventory)
	if err != nil {
		return nil, err
	}

	// Log the creation
	logDetails := map[string]interface{}{
		"inventory_id": inventory.ID,
		"name":         name,
		"company_id":   companyID,
		"category_id":  categoryID,
	}
	if err := repositories.CreateLog(userID, "Inventory", inventory.ID, "Created inventory item", logDetails); err != nil {
		return nil, err
	}

	return inventory, nil
}

// ArchiveInventory archives an inventory item
func ArchiveInventory(ctx context.Context, id int) error {
	userID, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.New("failed to retrieve user ID from context")
	}

	if id <= 0 {
		return errors.New("invalid inventory ID")
	}

	// Archive the inventory item
	if err := repositories.UpdateInventoryArchivedStatus(id, true); err != nil {
		return err
	}

	// Log the action
	logDetails := map[string]interface{}{"inventory_id": id, "action": "archived"}
	return repositories.CreateLog(userID, "Inventory", id, "Archived inventory item", logDetails)
}

// UnarchiveInventory unarchives an inventory item
func UnarchiveInventory(ctx context.Context, id int) error {
	userID, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.New("failed to retrieve user ID from context")
	}

	if id <= 0 {
		return errors.New("invalid inventory ID")
	}

	// Unarchive the inventory item
	if err := repositories.UpdateInventoryArchivedStatus(id, false); err != nil {
		return err
	}

	// Log the action
	logDetails := map[string]interface{}{"inventory_id": id, "action": "unarchived"}
	return repositories.CreateLog(userID, "Inventory", id, "Unarchived inventory item", logDetails)
}

// GetInventoryByID fetches an inventory item by its ID

func GetInventoryByID(ctx context.Context, id int) (*models.Inventory, error) {

	// Mock implementation, replace with actual database call

	if id <= 0 {

		return nil, errors.New("invalid inventory ID")

	}

	return &models.Inventory{

		ID: id,

		Name: "Sample Inventory",
	}, nil

}
