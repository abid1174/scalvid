package product

import (
	"encoding/json"
	"net/http"
	"scalvid/repo"
	"scalvid/utils"
	"strconv"
)

func (h *Handler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product repo.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedProduct, err := h.productRepo.UpdateProduct(id, product)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}
	utils.SendResponse(w, http.StatusOK, updatedProduct)
}
