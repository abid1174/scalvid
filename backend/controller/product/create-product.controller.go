package controller

import (
	"encoding/json"
	"net/http"
	"scalvid/data"
	"scalvid/entities"
	"scalvid/utils"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var products = data.Products
	var product entities.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	product.ID = len(products) + 1

	data.Products = append(data.Products, product)
	utils.SendResponse(w, http.StatusCreated, product)
}
