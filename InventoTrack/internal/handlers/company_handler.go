package handlers

import (
	"encoding/json"
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"net/http"
)

// CreateCompany handles POST /companies
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var companyRequest struct {
		Name string `json:"name"`
	}

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&companyRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Call the service to create the company
	company, err := services.CreateCompany(companyRequest.Name)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the created company
	utils.RespondWithJSON(w, http.StatusCreated, company)
}
