package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/api"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/dbase"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/internal"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/utils"
)

// Global validator used for json parsing
var validate *validator.Validate

func init() {
	//build validatorv10 json parser rules
	validate = validator.New()
	validate.RegisterValidation("gamegenre", internal.IsValidGameGenre)
}

func main() {
	utils.GetAppConfig() // instantiate app variables
	dbase.ConnectDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/games", api.GetGames).Methods("GET")
	router.HandleFunc("/games/{id}", api.GetGames).Methods("GET")
	router.HandleFunc("/games", api.CreateGame).Methods("POST")
	router.HandleFunc("/games/{id}", api.UpdateGame).Methods("PUT")
	router.HandleFunc("/games/{id}", api.DeleteGame).Methods("DELETE")

	router.HandleFunc("/boardgames", func(w http.ResponseWriter, r *http.Request) { api.GetBoardGames(w, r, dbase.DB) }).Methods("GET")
	router.HandleFunc("/boardgames/{id}", api.CreateBoardGame).Methods("POST")

	log.Println("Server listening on port 9999")
	log.Fatal(http.ListenAndServe(":"+utils.CFG.RestPort, router))
}
