package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"scalvid/infra/db"
	"scalvid/repo"
	"scalvid/utils"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateProductHandler")

	var product repo.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	log.Println("Invalid request body", err)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if product.Title == "" {
		utils.SendError(w, http.StatusBadRequest, "Title is required")
		return
	}
	if product.Price <= 0 {
		utils.SendError(w, http.StatusBadRequest, "Price must be greater than 0")
		return
	}

	dbConnection := db.NewConnection()

	productRepository := repo.NewProductRepository(dbConnection)
	createdProduct, err := productRepository.CreateProduct(repo.Product{
		Title:       product.Title,
		Price:       product.Price,
		Description: product.Description,
		Category:    product.Category,
		ImageUrl:    product.ImageUrl,
	})

	log.Println(createdProduct)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}
	utils.SendResponse(w, http.StatusCreated, createdProduct)
}
