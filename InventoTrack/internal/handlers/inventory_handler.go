package handlers

import (
	"encoding/json"
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"net/http"
	"strconv"
)

// AddInventory handles POST /inventory
func AddInventory(w http.ResponseWriter, r *http.Request) {
	var inventoryRequest struct {
		Name       string `json:"name"`
		CompanyID  int    `json:"company_id"`
		CategoryID int    `json:"category_id"`
	}

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&inventoryRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Call the service to add inventory
	inventory, err := services.AddInventory(inventoryRequest.Name, inventoryRequest.CompanyID, inventoryRequest.CategoryID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the created inventory item
	utils.RespondWithJSON(w, http.StatusCreated, inventory)
}

// GetInventory handles GET /inventory/{id}
func GetInventory(w http.ResponseWriter, r *http.Request) {
	idStr := utils.GetPathParam(r, "id") // Extract `id` from URL path
	id, err := strconv.Atoi(idStr)       // Convert `id` from string to int
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid inventory ID")
		return
	}

	// Call service to fetch the inventory
	inventory, err := services.GetInventoryByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Inventory item not found")
		return
	}

	// Respond with inventory data
	utils.RespondWithJSON(w, http.StatusOK, inventory)
}
