package handler

import (
	"encoding/json"
	"net/http"

	"product-service/model"
	"product-service/service"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	err := service.AddProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product added"})
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := service.ListProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
