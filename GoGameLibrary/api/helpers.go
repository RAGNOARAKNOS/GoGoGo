package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Extract the ID from the HTTP request, returns true if error occurs
func extractIdFromQuery(r *http.Request, w http.ResponseWriter) (int, bool) {
	vars := mux.Vars(r)
	idString := vars["id"]
	if idString == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return -1, true
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return -1, true
	}
	return id, false
}
