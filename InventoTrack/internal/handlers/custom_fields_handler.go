package handlers

import (
	"encoding/json"
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"net/http"
	"strconv"
)

// AddCustomField handles POST /custom-fields
func AddCustomField(w http.ResponseWriter, r *http.Request) {
	var request struct {
		TableName       string `json:"table_name"`
		FieldName       string `json:"field_name"`
		FieldType       string `json:"field_type"`
		ValidationRules string `json:"validation_rules"`
		CompanyID       int    `json:"company_id"`
	}

	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Call the service to add the custom field
	err := services.AddCustomField(request.TableName, request.FieldName, request.FieldType, request.ValidationRules, request.CompanyID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Custom field added successfully"})
}

// RemoveCustomField handles DELETE /custom-fields/{id}
func RemoveCustomField(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the path parameters
	idStr := utils.GetPathParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid custom field ID")
		return
	}

	// Call the service to remove the custom field
	err = services.RemoveCustomField(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Custom field removed successfully"})
}
