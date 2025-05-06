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

func GetTagsHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var Tags []dbase.Tag
	result := db.Find(&Tags)

	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to fetch Tags: %v", result.Error), http.StatusInternalServerError)
		return
	}

	TagResponses := make([]dtos.TagResponse, len(Tags))
	for i, m := range Tags {
		TagResponses[i] = dtos.TagResponse{
			ID:   m.ID,
			Name: m.Name,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TagResponses)
}

func GetTagHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
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
	var existingTag dbase.Tag
	result := db.First(&existingTag, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Tag not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to fetch Tag: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 3 Respond
	response := dtos.TagResponse{
		ID:   existingTag.ID,
		Name: existingTag.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreateTagHandler(w http.ResponseWriter, r *http.Request, v *validator.Validate, db *gorm.DB) {
	var req dtos.NewTagRequest

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
	Tag := dbase.Tag{
		Name: req.Name,
	}
	result := db.Create(&Tag)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to create Tag: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 4 Map GORM struct to DTO struct
	response := dtos.TagResponse{
		ID:   Tag.ID,
		Name: Tag.Name,
	}

	// 5 Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func UpdateTagHandler(w http.ResponseWriter, r *http.Request, v *validator.Validate, db *gorm.DB) {
	// 1 Decode the JSON request into DTO struct
	var req dtos.UpdateTagRequest
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

	// 4 Fetch the existing Tag from the dbase
	var existingTag dbase.Tag
	result := db.First(&existingTag, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Tag not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to fetch Tag: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 5 Update the fields
	if req.Name != "" {
		existingTag.Name = req.Name
	}

	// 6 Update the record in dbase
	result = db.Save(&existingTag)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("Failed to update Tag: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 7 Map GORM struct to DTO struct
	response := dtos.TagResponse{
		ID:   existingTag.ID,
		Name: existingTag.Name,
	}

	// 8 Send Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DeleteTagHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
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

	// 2 Fetch the Tag from the database
	var Tag dbase.Tag
	result := db.First(&Tag, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Tag not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("FAILED to fetch Tag: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 3 Delete the record from the database
	result = db.Delete(&Tag)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to delete Tag: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 4 Send the success response
	w.WriteHeader(http.StatusNoContent)
}
