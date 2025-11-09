package product

import (
	"encoding/json"
	"net/http"
	"scalvid/database"
	"scalvid/utils"
	"strconv"
)

func (h *Handler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL path parameter
	idStr := r.PathValue("id")
	if idStr == "" {
		utils.SendError(w, http.StatusBadRequest, "Product ID is required")
		return
	}

	// Convert ID to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	// Check if product exists
	existingProduct := database.GetProduct(id)
	if existingProduct.ID == 0 {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	// Decode request body
	var product database.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Set the ID from URL (prevent ID change)
	product.ID = id

	// Update product
	database.UpdateProduct(id, product)

	utils.SendResponse(w, http.StatusOK, product)
}
