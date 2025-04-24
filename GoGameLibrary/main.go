package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/api"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/dbase"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/utils"
)

func main() {
	utils.GetAppConfig() // instantiate app variables
	dbase.ConnectDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/games", api.GetGames).Methods("GET")
	router.HandleFunc("/games/{id}", api.GetGames).Methods("GET")
	router.HandleFunc("/games", api.CreateGame).Methods("POST")
	router.HandleFunc("/games/{id}", api.UpdateGame).Methods("PUT")
	router.HandleFunc("/games/{id}", api.DeleteGame).Methods("DELETE")

	router.HandleFunc("/boardgames", api.GetBoardGames).Methods("GET")

	log.Println("Server listening on port 9999")
	log.Fatal(http.ListenAndServe(":"+utils.CFG.RestPort, router))
}
