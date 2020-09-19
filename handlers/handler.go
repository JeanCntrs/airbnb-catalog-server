package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// RouteHandlers : Raise the server on a specific port, giving permission to any client to consult
func RouteHandlers() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	corsHandler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, corsHandler))
}
