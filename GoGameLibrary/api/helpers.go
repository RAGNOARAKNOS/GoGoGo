package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
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

// Extract the ID from URL path variables
// Better than my first attempt extractIdFromQuery, adds better validation
func parseID(r *http.Request, paramName string) (uint, error) {
	vars := mux.Vars(r)
	idStr, ok := vars[paramName]
	if !ok {
		return 0, fmt.Errorf("missing path parameter: %s", paramName)
	}
	id, err := strconv.ParseUint(idStr, 10, 32) //Parse UInt to match ID in struct
	if err != nil || id == 0 {                  // Validate that ID is positive
		return 0, fmt.Errorf("invalid %s ID: %s", paramName, idStr)
	}
	return uint(id), nil
}

// Sends a JSON response with a given status code and payload
// "any" is an alias for interface{}, i.e an empty interface definition
func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			// Log the marshalling error, return a generic server error
			log.Printf("ERROR: Failed to marshal JSON response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Internal server error"}`))
			return
		}
		w.WriteHeader(status)
		w.Write(response)
	} else {
		w.WriteHeader(status) // write header for nil payload (e.g. 204 no content)
	}
}

// Sends a JSON error response
func respondError(w http.ResponseWriter, code int, message string) {
	log.Printf("ERROR %d: %s", code, message)
	respondJSON(w, code, map[string]string{"error": message})
}

// Parses validator errors and sends structured error response
func handleValidationErrors(w http.ResponseWriter, err error) {
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		// Build a map of fields -> error message
		errorMessages := make(map[string]string)
		for _, fe := range verr {
			field := fe.Error()
			// Build more user-friendly messages based on the validation tag
			switch fe.Tag() {
			case "required":
				errorMessages[field] = fmt.Sprintf("%s is required", field)
			case "min":
				errorMessages[field] = fmt.Sprintf("%s must be at least %s characters long", field, fe.Param())
			case "max":
				errorMessages[field] = fmt.Sprintf("%s must be at most %s characters long", field, fe.Param())
			case "gte":
				errorMessages[field] = fmt.Sprintf("%s must be greater than or equal to %s", field, fe.Param())
			case "gt":
				errorMessages[field] = fmt.Sprintf("%s must be greater than %s", field, fe.Param())
			case "gtefield":
				errorMessages[field] = fmt.Sprintf("%s must be greater than or equal to %s", field, fe.Param())
			case "unique":
				errorMessages[field] = fmt.Sprintf("%s must contain unique values", field)
			case "dive":
				// Error is within a slice/map, CBA to iterate down, this should suffice
				errorMessages[field] = fmt.Sprintf("Invalid value found within %s", field)
			default:
				errorMessages[field] = fmt.Sprintf("Invalid value for %s (failed on '%s' validation)", field, fe.Tag())
			}
		}
		log.Printf("Validation Errors: %v", errorMessages)
		respondJSON(w, http.StatusBadRequest, map[string]interface{}{
			"error":             "Validation Failed",
			"validation_errors": errorMessages,
		})
		return
	}
	//Handle non-validation errors (e.g. JSON decode issues)
	log.Printf("ERROR: Bad Request (non-validation): %v", err)
	respondError(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
}
