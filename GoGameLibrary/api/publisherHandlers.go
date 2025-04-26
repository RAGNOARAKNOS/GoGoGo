package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/dbase"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/dtos"
	"gorm.io/gorm"
)

func GetPublishersHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var publishers []dbase.Publisher
	result := db.Find(&publishers)

	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to fetch Publishers: %v", result.Error), http.StatusInternalServerError)
		return
	}

	publisherResponses := make([]dtos.PublisherResponse, len(publishers))
	for i, p := range publishers {
		publisherResponses[i] = dtos.PublisherResponse{
			ID:   p.ID,
			Name: p.Name,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(publisherResponses)
}

func GetPublisherHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// 1 Get the Id from the request
	id, err := extractIdFromQuery(r, w)
	if err {
		return
	}

	// 2 Fetch from DBASE
	var existingPublisher dbase.Publisher
	result := db.First(&existingPublisher, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Publisher not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("FAILED to fetch publisher: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 3 Respond
	response := dtos.PublisherResponse{
		ID:   existingPublisher.ID,
		Name: existingPublisher.Name,
		//Games: []dtos.BoardgameResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreatePublisherHandler(w http.ResponseWriter, r *http.Request, v *validator.Validate, db *gorm.DB) {
	var req dtos.NewPublisherRequest

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

	// 3 Create new dbase record
	publisher := dbase.Publisher{
		Name: req.Name,
	}

	result := db.Create(&publisher)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to create publisher: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 4 Map GORM struct to DTO struct
	response := dtos.PublisherResponse{
		ID:   publisher.ID,
		Name: publisher.Name,
	}

	// 5 Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func UpdatePublisherHandler(w http.ResponseWriter, r *http.Request, v *validator.Validate, db *gorm.DB) {
	// 1 Decode the JSON request into DTO struct
	var req dtos.PublisherResponse

	// 2 Validate the request struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	// 3 Get the ID from the request
	id, err := extractIdFromQuery(r, w)
	if err {
		return
	}

	// 4 Fetch the existing publisher from the dbase
	var existingPublisher dbase.Publisher
	result := db.First(&existingPublisher, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Publisher not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to fetch publisher: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 5 Update the fields
	if req.Name != "" {
		existingPublisher.Name = req.Name
	}

	// 6 Update the record in the dbase
	result = db.Save(&existingPublisher)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("Failed to update publisher: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 7 Map GORM struct to DTO struct
	response := dtos.PublisherResponse{
		ID:   existingPublisher.ID,
		Name: existingPublisher.Name,
	}

	// 8 Send Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DeletePublisherHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// 1 Get Id from request
	id, err := extractIdFromQuery(r, w)
	if err {
		return
	}

	// 2 Fetch the publisher from the database
	var publisher dbase.Publisher
	result := db.First(&publisher, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Publisher not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("FAILED to fetch publisher: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 3 Delete the record from the dbase
	result = db.Delete(&publisher)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to delete publisher: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 4 Send the success response
	w.WriteHeader(http.StatusNoContent)
}
