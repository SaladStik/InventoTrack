package handlers

import (
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"net/http"
	"strconv"
)

// GetRecycleBin handles GET /recycle-bin
func GetRecycleBin(w http.ResponseWriter, r *http.Request) {
	companyID := utils.GetCompanyIDFromContext(r.Context())
	if companyID == 0 {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or missing company ID")
		return
	}

	entries, err := services.GetRecycleBinEntries(companyID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, entries)
}

// PermanentlyDelete handles DELETE /recycle-bin/{id}
func PermanentlyDelete(w http.ResponseWriter, r *http.Request) {
	idStr := utils.GetPathParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid recycle bin entry ID")
		return
	}

	err = services.PermanentlyDeleteEntry(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Entry permanently deleted"})
}
