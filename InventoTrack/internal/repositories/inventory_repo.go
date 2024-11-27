package repositories

import (
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
)

// CreateInventory adds a new inventory item to the database
func CreateInventory(inventory *models.Inventory) error {
	result := extensions.DB.Create(inventory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetInventoryByID retrieves an inventory item by its ID
func GetInventoryByID(id int) (*models.Inventory, error) {
	var inventory models.Inventory
	result := extensions.DB.First(&inventory, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &inventory, nil
}
