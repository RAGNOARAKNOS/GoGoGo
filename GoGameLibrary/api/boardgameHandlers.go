package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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
			Genre:       bg.Genre,
			Complexity:  bg.Complexity,
			MinPlayers:  bg.MinPlayers,
			MaxPlayers:  bg.MaxPlayers,
			BestPlayers: bg.BestPlayers,
			Playtime:    bg.Playtime,
			Designer:    bg.Designer,
			PublisherID: bg.PublisherID,
			ImageURL:    bg.ImageURL,
			ParentID:    int(*bg.ParentID),
		}
	}

	// 3. Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(boardgameResponses)
}
