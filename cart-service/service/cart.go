package service

import (
	"cart-service/model"
)

var cartItems = []model.CartItem{}

func AddToCart(item model.CartItem) error {
	cartItems = append(cartItems, item)
	return nil
}

func ViewCart() []model.CartItem {
	return cartItems
}
