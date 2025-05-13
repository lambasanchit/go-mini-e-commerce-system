package handler

import (
	"encoding/json"
	"net/http"

	"cart-service/model"
	"cart-service/service"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	var item model.CartItem
	json.NewDecoder(r.Body).Decode(&item)

	err := service.AddToCart(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item added to cart"})
}

func ViewCart(w http.ResponseWriter, r *http.Request) {
	cart := service.ViewCart()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}
