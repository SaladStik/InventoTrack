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
		CompanyID       int    `json:"company_id"`
		TableName       string `json:"table_name"`
		FieldName       string `json:"field_name"`
		FieldType       string `json:"field_type"`
		ValidationRules string `json:"validation_rules"`
	}

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Call the service
	if err := services.AddCustomField(request.CompanyID, request.TableName, request.FieldName, request.FieldType, request.ValidationRules); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Custom field added successfully"})
}

// RemoveCustomField handles DELETE /custom-fields/{id}
func RemoveCustomField(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(utils.GetPathParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid custom field ID")
		return
	}

	if err := services.RemoveCustomField(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Custom field removed successfully"})
}

// AddCustomFieldValue handles POST /custom-field-values
func AddCustomFieldValue(w http.ResponseWriter, r *http.Request) {
	var request struct {
		CustomFieldID int    `json:"custom_field_id"`
		InventoryID   int    `json:"inventory_id"`
		Value         string `json:"value"`
	}

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Call the service
	if err := services.AddCustomFieldValue(request.CustomFieldID, request.InventoryID, request.Value); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Custom field value added successfully"})
}

// GetCustomFieldValues handles GET /custom-field-values/{entityID}
func GetCustomFieldValues(w http.ResponseWriter, r *http.Request) {
	entityID, err := strconv.Atoi(utils.GetPathParam(r, "entityID"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid entity ID")
		return
	}

	values, err := services.GetCustomFieldValues(entityID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, values)
}
