package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

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
		respondError(w, http.StatusBadRequest, "Invalid request payload:"+err.Error())
		return
	}
	defer r.Body.Close()

	// 2 Validate the request
	if err := v.Struct(req); err != nil {
		handleValidationErrors(w, err)
		return
	}

	// 3 Create new dbase record
	// 3a Insert the basic data
	game := dbase.Boardgame{
		Title:        req.Title,
		Description:  req.Description,
		ImageURL:     req.ImageURL,
		Complexity:   req.Complexity,
		Learnability: req.Learnability,
		Playtime:     req.Playtime,
		Setuptime:    req.Setuptime,
		MinPlayers:   req.MinPlayers,
		MaxPlayers:   req.MaxPlayers,
		BestPlayers:  req.BestPlayers,
		Designer:     req.Designer,
		//BGGURL:       req.BoardGameGeekURL,
		//BGGRating:    req.BoardGameGeekRating,
		PublisherID: req.PublisherID,
		BasegameID:  req.BasegameID, // assign the pointer to the base game
	}

	// 3b insert the relationship-based data
	// done within a 'transaction' to make rollback cleaner if the process fails
	err := db.Transaction(func(tx *gorm.DB) error {
		// 3b-1 Validate the Publisher
		if req.PublisherID != nil && *req.PublisherID > 0 {
			var pub dbase.Publisher
			if err := tx.Select("id").First(&pub, *req.PublisherID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("publisher ID %d not found", *req.PublisherID)
				}
			} else if req.PublisherID != nil && *req.PublisherID == 0 {
				return fmt.Errorf("publisher id cannot be 0")
			}
		}
		// 3b-2 Validate the BaseGame
		if req.BasegameID != nil && *req.BasegameID > 0 {
			var baseGame dbase.Boardgame
			if err := tx.Select("id").First(&baseGame, *req.BasegameID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("base game ID %d not found", *req.BasegameID) // Return specific error
				}
				return err // Other DB error
			} else if req.BasegameID != nil && *req.BasegameID == 0 {
				return fmt.Errorf("base game ID cannot be 0")
			}
		}
		// 3b-3 Lookup the Genres
		if len(req.GenreIDs) > 0 {
			var genres []*dbase.Tag
			if err := tx.Where("id in ?", req.GenreIDs).Find(&genres).Error; err != nil {
				return fmt.Errorf("failed to query genres: %w", err)
			}
			if len(genres) != len(req.GenreIDs) { // genre mismatch in list
				return fmt.Errorf("one or more specified genre IDs not found")
			}
			game.Genres = genres // Assign the genre tags
		}
		// 3b-4 Lookup the Mechanics
		if len(req.MechanicIDs) > 0 {
			var mechanics []*dbase.Tag
			if err := tx.Where("id in ?", req.MechanicIDs).Find(&mechanics).Error; err != nil {
				return fmt.Errorf("failed to query mechanics: %w", err)
			}
			if len(mechanics) != len(req.MechanicIDs) {
				return fmt.Errorf("one or more specified mechanic IDs not found")
			}
			game.Mechanics = mechanics // Assign the mechanic tags
		}
		// 3b-5 Lookup the Bits
		if len(req.BitIDs) > 0 {
			var bits []*dbase.Tag
			if err := tx.Where("id in ?", req.BitIDs).Find(&bits).Error; err != nil {
				return fmt.Errorf("failed to query bits: %w", err)
			}
			if len(bits) != len(req.BitIDs) {
				return fmt.Errorf("one or more specified Bit IDs not found")
			}
			game.Bits = bits // Assign the but tags
		}

		// 3c Create the BoardGame record
		if err := tx.Create(&game).Error; err != nil {
			return err // return creation error
		}
		return nil // commit transaction
	}) // END OF TRANSACTION

	// Handle potential transaction errors (e.g tag not found, base game not found, db errors)
	if err != nil {
		// Check if a *specific* validation error was returned
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "cannot be 0") {
			respondError(w, http.StatusBadRequest, err.Error())
		} else {
			handleGormError(w, err, "definition creation error") // Assume server error
		}
		return
	}

	// 4 Map GORM struct to DTO struct
	if err := db.
		Preload("Publisher").
		Preload("Genres").
		Preload("Mechanics").
		Preload("Expansions").
		First(&game, game.ID).Error; err != nil {
		log.Printf("WARN: Failed to reload definition %d after create: %v", game.ID, err)
	}

	// 5 Send success response
	respondJSON(w, http.StatusCreated, mapModelToResponse(&game))

}

func mapModelToResponse(model *dbase.Boardgame) dtos.BoardgameResponse {
	if model == nil {
		return dtos.BoardgameResponse{}
	}

	expansionsSummary := make([]dtos.BoardgameTerseResponse, 0, len(model.Expansions))
	for _, exp := range model.Expansions {
		if exp != nil {
			expansionsSummary = append(expansionsSummary, mapModelToTerseResponse(exp))
		}
	}

	return dtos.BoardgameResponse{
		ID:           model.ID,
		Title:        model.Title,
		Description:  model.Description,
		Complexity:   model.Complexity,
		Learnability: model.Learnability,
		MinPlayers:   model.MinPlayers,
		MaxPlayers:   model.MaxPlayers,
		BestPlayers:  model.BestPlayers,
		Playtime:     model.BestPlayers,
		Designer:     model.Designer,
		Expansions:   expansionsSummary,
	}
}

// Map a SHORT game summary response from the data model
func mapModelToTerseResponse(model *dbase.Boardgame) dtos.BoardgameTerseResponse {
	if model == nil {
		return dtos.BoardgameTerseResponse{}
	}
	return dtos.BoardgameTerseResponse{
		ID:    model.ID,
		Title: model.Title,
	}
}

// Map a SHORT summary response from the data model for each entry
func mapModelsToTerseResponse(models []*dbase.Boardgame) []dtos.BoardgameTerseResponse {
	resp := make([]dtos.BoardgameTerseResponse, 0, len(models))
	for _, d := range models {
		if d != nil {
			resp = append(resp, mapModelToTerseResponse(d))
		}
	}
	return resp
}
