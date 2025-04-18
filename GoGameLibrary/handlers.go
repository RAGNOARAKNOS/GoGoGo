package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetGames(w http.ResponseWriter, r *http.Request) {
	var games []Game
	DB.Find(&games)
	json.NewEncoder(w).Encode(games)
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	var game Game
	result := DB.First(&game, id)
	if result.Error != nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(game)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var game Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := DB.Create(&game)
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

	var game Game
	result := DB.First(&game, id)
	if result.Error != nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	var updatedGame Game
	err = json.NewDecoder(r.Body).Decode(&updatedGame)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	game.Title = updatedGame.Title
	game.Studio = updatedGame.Studio
	game.Publisher = updatedGame.Publisher

	DB.Save(&game)
	json.NewEncoder(w).Encode(game)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	var game Game
	result := DB.Delete(&game, id)
	if result.Error != nil {
		http.Error(w, "Failed to delete game", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
