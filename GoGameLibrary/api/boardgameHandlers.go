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

func GetBoardGames(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// 1. Get boardgames from DB, ensuring all linkages are preloaded
	var boardGames []dbase.Boardgame
	result := db.Preload("Publisher").Preload("Mechanics").Preload("Parent").Preload("Expansions").Find(&boardGames)

	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to fetch board games: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 2. Map the GORM struct to DTO
	boardgameResponses := make([]dtos.BoardgameResponse, len(boardGames))
	for i, bg := range boardGames {
		boardgameResponses[i] = dtos.BoardgameResponse{
			ID:          bg.ID,
			Title:       bg.Title,
			Description: bg.Description,
			//Genre:       bg.Genre,
			Complexity:  bg.Complexity,
			MinPlayers:  bg.MinPlayers,
			MaxPlayers:  bg.MaxPlayers,
			BestPlayers: bg.BestPlayers,
			Playtime:    bg.Playtime,
			Designer:    bg.Designer,
			//PublisherID: bg.PublisherID,
			ImageURL: bg.ImageURL,
			ParentID: int(*bg.BasegameID),
		}
	}

	// 3. Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(boardgameResponses)
}

func CreateBoardgameHandler(w http.ResponseWriter, r *http.Request, v *validator.Validate, db *gorm.DB) {
	var req dtos.NewBoardgameRequest

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
	game := dbase.Boardgame{
		Title:       req.Title,
		Description: req.Description,
		//Genre:       req.Genre,
		Complexity:  req.Complexity,
		MinPlayers:  req.MinPlayers,
		MaxPlayers:  req.MaxPlayers,
		BestPlayers: req.BestPlayers,
		Playtime:    req.Playtime,
		Designer:    req.Designer,
		PublisherID: req.PublisherID,
		ImageURL:    req.ImageURL,
	}

	result := db.Create(&game)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("FAILED to create board game: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// 4 Map GORM struct to DTO struct
	response := dtos.BoardgameResponse{
		ID:          game.ID,
		Title:       game.Title,
		Description: game.Description,
	}

	// 5 Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
