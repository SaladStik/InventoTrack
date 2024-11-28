package repositories

import (
	"database/sql"
	"errors"
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
	"log"
	"time"
)

type InventoryRepo struct {
	db *sql.DB
}

type InventoryItem struct {
	CompanyID       int
	Name            string
	CategoryID      int
	IsArchived      bool
	RetentionPeriod int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// CreateInventory adds a new inventory item to the database
func CreateInventory(inventory *models.Inventory) error {
	if inventory.Name == "" {
		return errors.New("inventory name cannot be empty")
	}
	if inventory.CompanyID <= 0 {
		return errors.New("invalid company ID")
	}

	// Attempt to create the inventory item
	result := extensions.DB.Create(inventory)
	if result.Error != nil {
		log.Printf("Failed to create inventory: %v", result.Error)
		return result.Error
	}

	log.Printf("Inventory item created: %v", inventory)
	return nil
}

// GetInventoryByID retrieves an inventory item by its ID
func GetInventoryByID(id int) (*models.Inventory, error) {
	if id <= 0 {
		return nil, errors.New("invalid inventory ID")
	}

	var inventory models.Inventory
	result := extensions.DB.First(&inventory, id)
	if result.Error != nil {
		log.Printf("Failed to retrieve inventory item with ID %d: %v", id, result.Error)
		return nil, result.Error
	}

	log.Printf("Retrieved inventory item: %v", inventory)
	return &inventory, nil
}

// UpdateInventoryArchivedStatus updates the archived status of an inventory item
func UpdateInventoryArchivedStatus(id int, isArchived bool) error {
	if id <= 0 {
		return errors.New("invalid inventory ID")
	}

	// Update the archived status
	result := extensions.DB.Model(&models.Inventory{}).Where("id = ?", id).Update("is_archived", isArchived)
	if result.Error != nil {
		log.Printf("Failed to update archive status for inventory ID %d: %v", id, result.Error)
		return result.Error
	}

	log.Printf("Updated archive status for inventory ID %d to %v", id, isArchived)
	return nil
}

func (repo *InventoryRepo) CreateInventoryItem(item InventoryItem, userID int) (int, error) {
	var inventoryID int
	err := repo.db.QueryRow(`
		INSERT INTO "inventory" ("company_id","name","category_id","is_archived","retention_period","created_at","updated_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING "id"`,
		item.CompanyID, item.Name, item.CategoryID, item.IsArchived, item.RetentionPeriod, item.CreatedAt, item.UpdatedAt).Scan(&inventoryID)
	if err != nil {
		return 0, err
	}

	// Ensure userID is not null
	if userID == 0 {
		return 0, errors.New("userID cannot be null")
	}

	// Insert log entry with user_id
	_, err = repo.db.Exec(`
		INSERT INTO "logs" ("user_id", "action", "timestamp")
		VALUES ($1, $2, $3)`,
		userID, "Created inventory item", time.Now())
	if err != nil {
		return 0, errors.New("failed to create log entry: " + err.Error())
	}

	return inventoryID, nil
}
