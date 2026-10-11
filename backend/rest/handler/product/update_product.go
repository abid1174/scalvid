package product

import (
	"errors"
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
	err = utils.DecodeJSON(w, r, &product)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validateProduct(product); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedProduct, err := h.productRepo.UpdateProduct(r.Context(), id, product)
	if errors.Is(err, repo.ErrNotFound) {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}
	utils.SendResponse(w, http.StatusOK, updatedProduct)
}
