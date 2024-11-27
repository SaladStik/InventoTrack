package utils

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GetPathParam extracts a path parameter from the request URL
func GetPathParam(r *http.Request, key string) string {
	// First, check if Gorilla Mux is being used
	vars := mux.Vars(r)
	if val, ok := vars[key]; ok {
		return val
	}

	// Fallback for manual URL parsing
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if part == key && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}
