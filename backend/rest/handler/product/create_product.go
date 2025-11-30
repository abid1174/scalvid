package product

import (
	"encoding/json"
	"log"
	"net/http"
	"scalvid/repo"
	"scalvid/utils"
)

func (h *Handler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateProductHandler")

	var product repo.Product

	log.Println("Request Body", r.Body)

	err := json.NewDecoder(r.Body).Decode(&product)
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

	createdProduct, err := h.productRepo.CreateProduct(repo.Product{
		Title:       product.Title,
		Price:       product.Price,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
	})

	log.Println(createdProduct)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}
	utils.SendResponse(w, http.StatusCreated, createdProduct)
}
