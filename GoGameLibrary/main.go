package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/api"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/dbase"
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/internal"
)

// Global validator used for json parsing
var validate *validator.Validate

func init() {
	//build validatorv10 json parser rules
	validate = validator.New()
	validate.RegisterValidation("gamegenre", internal.IsValidGameGenre)
}

func main() {
	internal.GetAppConfig() // instantiate app variables
	dbase.ConnectDatabase()

	router := mux.NewRouter()

	// Deprecated API
	router.HandleFunc("/games", api.GetGames).Methods("GET")
	router.HandleFunc("/games/{id}", api.GetGames).Methods("GET")
	router.HandleFunc("/games", api.CreateGame).Methods("POST")
	router.HandleFunc("/games/{id}", api.UpdateGame).Methods("PUT")
	router.HandleFunc("/games/{id}", api.DeleteGame).Methods("DELETE")

	// Board Game API endpoints
	router.HandleFunc("/boardgames", func(w http.ResponseWriter, r *http.Request) { api.GetBoardGames(w, r, dbase.DB) }).Methods("GET")
	router.HandleFunc("/boardgames/{id}", api.CreateBoardGame).Methods("POST")

	// Tag API endpoints
	router.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) { api.GetTagsHandler(w, r, dbase.DB) }).Methods("GET")
	router.HandleFunc("/tags/{id}", func(w http.ResponseWriter, r *http.Request) { api.GetTagHandler(w, r, dbase.DB) }).Methods("GET")
	router.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) { api.CreateTagHandler(w, r, validate, dbase.DB) }).Methods("POST")
	router.HandleFunc("/tags/{id}", func(w http.ResponseWriter, r *http.Request) { api.UpdateTagHandler(w, r, validate, dbase.DB) }).Methods("PATCH")
	router.HandleFunc("/tags/{id}", func(w http.ResponseWriter, r *http.Request) { api.DeleteTagHandler(w, r, dbase.DB) }).Methods("DELETE")

	// Publisher API endpoints
	router.HandleFunc("/publishers", func(w http.ResponseWriter, r *http.Request) { api.GetPublishersHandler(w, r, dbase.DB) }).Methods("GET")
	router.HandleFunc("/publishers/{id}", func(w http.ResponseWriter, r *http.Request) { api.GetPublisherHandler(w, r, dbase.DB) }).Methods("GET")
	router.HandleFunc("/publishers", func(w http.ResponseWriter, r *http.Request) { api.CreatePublisherHandler(w, r, validate, dbase.DB) }).Methods("POST")
	router.HandleFunc("/publishers/{id}", func(w http.ResponseWriter, r *http.Request) { api.UpdatePublisherHandler(w, r, validate, dbase.DB) }).Methods("PATCH")
	router.HandleFunc("/publishers/{id}", func(w http.ResponseWriter, r *http.Request) { api.DeletePublisherHandler(w, r, dbase.DB) }).Methods("DELETE")

	// Server configuration
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + internal.CFG.RestPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server goroutine
	go func() {
		log.Printf("Server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("FATAL: ListenAndServer Error: %v", err)
		}
	}()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // ctrl+c or term failure signals closure

	sig := <-quit
	log.Printf("RECEIVED SIGNAL: %s. Shutting down server...", sig)

	// Create a context with 30 sec timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("FATAL: Server forced to shutdown: %v", err)
	}

	log.Println("Server Exiting.")
}
