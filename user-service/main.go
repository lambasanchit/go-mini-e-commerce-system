package main

import (
	"log"
	"net/http"
	"os"
	"user-service/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")

	// Use an environment variable for port flexibility
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" // Default to 8001 if not set
	}

	log.Printf("User Service running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
