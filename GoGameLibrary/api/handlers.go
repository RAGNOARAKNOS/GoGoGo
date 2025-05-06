package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/dbase"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	var games []dbase.Game
	dbase.DB.Find(&games)
	json.NewEncoder(w).Encode(games)
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	var game dbase.Game
	result := dbase.DB.First(&game, id)
	if result.Error != nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(game)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var game dbase.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := dbase.DB.Create(&game)
	if result.Error != nil {
		http.Error(w, "Failed to create game", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(game)
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	var game dbase.Game
	result := dbase.DB.First(&game, id)
	if result.Error != nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	var updatedGame dbase.Game
	err = json.NewDecoder(r.Body).Decode(&updatedGame)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	game.Title = updatedGame.Title
	game.Studio = updatedGame.Studio
	game.Publisher = updatedGame.Publisher

	dbase.DB.Save(&game)
	json.NewEncoder(w).Encode(game)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	var game dbase.Game
	result := dbase.DB.Delete(&game, id)
	if result.Error != nil {
		http.Error(w, "Failed to delete game", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// func GetBoardGames(w http.ResponseWriter, r *http.Request) {
// 	var games []dbase.Boardgame
// 	result := dbase.DB.Preload("Publisher").Preload("Mechanics").Preload("Expansions").Find(&games)

// 	if result.Error != nil { // handle database connection-style issues
// 		http.Error(w, "FAILED to fetch board games", http.StatusInternalServerError)
// 		return
// 	}

// 	if result.RowsAffected == 0 { //If no games are found
// 		json.NewEncoder(w).Encode([]dbase.Boardgame{}) //retutn empty array
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json") // return JSON
// 	json.NewEncoder(w).Encode(games)
// }

func GetBoardGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "INVALID board game ID", http.StatusBadRequest)
		return
	}

	var game dbase.Boardgame
	result := dbase.DB.Preload("Publisher").Preload("Mechanics").Preload("Expansions").First(&game, id)
	if result.Error != nil {
		http.Error(w, "Board game not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(game)
}

// sendErrorResponse sends a JSON error response
func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
