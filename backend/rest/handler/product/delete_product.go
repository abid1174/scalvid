package handler

import (
	"net/http"
	"scalvid/database"
	"scalvid/utils"
	"strconv"
)

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
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

	// Delete product
	database.DeleteProduct(id)

	utils.SendResponse(w, http.StatusOK, map[string]string{
		"message": "Product deleted successfully",
	})
}
