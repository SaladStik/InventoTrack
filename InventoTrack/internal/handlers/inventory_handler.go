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
		CategoryID *int   `json:"category_id"`
	}

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&inventoryRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Call the service to add inventory
	inventory, err := services.AddInventory(r.Context(), inventoryRequest.Name, inventoryRequest.CompanyID, inventoryRequest.CategoryID)
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
	inventory, err := services.GetInventoryByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Inventory item not found")
		return
	}

	// Respond with inventory data
	utils.RespondWithJSON(w, http.StatusOK, inventory)
}

// ArchiveInventory handles POST /inventory/{id}/archive
func ArchiveInventory(w http.ResponseWriter, r *http.Request) {
	idStr := utils.GetPathParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid inventory ID")
		return
	}

	// Pass the request context
	err = services.ArchiveInventory(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Inventory item archived"})
}

// UnarchiveInventory handles POST /inventory/{id}/unarchive
func UnarchiveInventory(w http.ResponseWriter, r *http.Request) {
	idStr := utils.GetPathParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid inventory ID")
		return
	}

	// Pass the request context
	err = services.UnarchiveInventory(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Inventory item unarchived"})
}
