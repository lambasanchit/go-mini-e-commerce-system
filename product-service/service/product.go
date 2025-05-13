package service

import (
	"errors"

	"product-service/model"
)

var products = []model.Product{}

func AddProduct(product model.Product) error {
	// Check if product already exists
	for _, p := range products {
		if p.Name == product.Name {
			return errors.New("product already exists")
		}
	}
	products = append(products, product)
	return nil
}

func ListProducts() []model.Product {
	return products
}
