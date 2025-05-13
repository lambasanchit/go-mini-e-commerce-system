package main

import (
	"log"
	"net/http"

	"cart-service/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/add", handler.AddToCart).Methods("POST")
	r.HandleFunc("/view", handler.ViewCart).Methods("GET")

	log.Println("Cart Service running on port 8003...")
	log.Fatal(http.ListenAndServe(":8003", r))
}
