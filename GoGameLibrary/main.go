package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	GetAppConfig() // instantiate app variables
	ConnectDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/games", GetGames).Methods("GET")
	router.HandleFunc("/games/{id}", GetGames).Methods("GET")
	router.HandleFunc("/games", CreateGame).Methods("POST")
	router.HandleFunc("/games/{id}", UpdateGame).Methods("PUT")
	router.HandleFunc("/games/{id}", DeleteGame).Methods("DELETE")

	log.Println("Server listening on port 9999")
	log.Fatal(http.ListenAndServe(":"+CFG.RestPort, router))
}
