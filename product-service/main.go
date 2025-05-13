package main

import (
	"log"
	"net/http"

	"product-service/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/add", handler.AddProduct).Methods("POST")
	r.HandleFunc("/list", handler.ListProducts).Methods("GET")

	log.Println("Product Service running on port 8002...")
	log.Fatal(http.ListenAndServe(":8002", r))
}
