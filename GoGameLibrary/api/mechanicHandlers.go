package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/dbase"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/dtos"
	"gorm.io/gorm"
)

func GetMechanicsHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var mechanics []dbase.Mechanic
	result := db.Find(&mechanics)

	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to fetch mechanics: %v", result.Error), http.StatusInternalServerError)
		return
	}

	mechanicResponses := make([]dtos.MechanicResponse, len(mechanics))
	for i, m := range mechanics {
		mechanicResponses[i] = dtos.MechanicResponse{
			ID:   m.ID,
			Name: m.Name,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mechanicResponses)
}

func GetMechanicHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// 1 Get the Id from the request
	vars := mux.Vars(r)
	idString := vars["id"]
	if idString == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// 2 Fetch from DBASE
	var existingMechanic dbase.Mechanic
	result := db.First(&existingMechanic, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Mechanic not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to fetch mechanic: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 3 Respond
	response := dtos.MechanicResponse{
		ID:   existingMechanic.ID,
		Name: existingMechanic.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreateMechanicHandler(w http.ResponseWriter, r *http.Request, v *validator.Validate, db *gorm.DB) {
	var req dtos.NewMechanicRequest

	// 1 Decode the JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 2 Validate the request
	if err := v.Struct(req); err != nil {
		errors := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errors = append(errors, fmt.Sprintf("Field %s is required", err.Field()))
			default:
				errors = append(errors, fmt.Sprintf("Field %s is invalid", err.Field()))
			}
		}
		errorResponse := map[string][]string{"errors": errors}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// 3 Create new DBASE record
	mechanic := dbase.Mechanic{
		Name: req.Name,
	}
	result := db.Create(&mechanic)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to create mechanic: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 4 Map GORM struct to DTO struct
	response := dtos.MechanicResponse{
		ID:   mechanic.ID,
		Name: mechanic.Name,
	}

	// 5 Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func UpdateMechanicHandler(w http.ResponseWriter, r *http.Request, v *validator.Validate, db *gorm.DB) {
	// 1 Decode the JSON request into DTO struct
	var req dtos.UpdateMechanicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "INVALID request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 2 Validate the request struct
	if err := v.Struct(req); err != nil {
		// Handle validation errors
		errors := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errors = append(errors, fmt.Sprintf("Field %s is required", err.Field()))
			default:
				errors = append(errors, fmt.Sprintf("Field %s is invalid", err.Field()))
			}
		}
		errorResponse := map[string][]string{"errors": errors}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// 3 Get the ID from request
	vars := mux.Vars(r)
	idString := vars["id"]
	if idString == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// 4 Fetch the existing mechanic from the dbase
	var existingMechanic dbase.Mechanic
	result := db.First(&existingMechanic, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Mechanic not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to fetch mechanic: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 5 Update the fields
	if req.Name != "" {
		existingMechanic.Name = req.Name
	}

	// 6 Update the record in dbase
	result = db.Save(&existingMechanic)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("Failed to update mechanic: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 7 Map GORM struct to DTO struct
	response := dtos.MechanicResponse{
		ID:   existingMechanic.ID,
		Name: existingMechanic.Name,
	}

	// 8 Send Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DeleteMechanicHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// 1 Get the ID from request
	vars := mux.Vars(r)
	idString := vars["id"]
	if idString == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// 2 Fetch the mechanic from the database
	var mechanic dbase.Mechanic
	result := db.First(&mechanic, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Mechanic not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("FAILED to fetch mechanic: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 3 Delete the record from the database
	result = db.Delete(&mechanic)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to delete mechanic: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 4 Send the success response
	w.WriteHeader(http.StatusNoContent)
}
