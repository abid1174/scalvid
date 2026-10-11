package product

import (
	"errors"
	"net/http"
	"scalvid/repo"
	"scalvid/utils"
	"strconv"
)

func (h *Handler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.productRepo.GetProduct(r.Context(), id)
	if errors.Is(err, repo.ErrNotFound) {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to get product")
		return
	}
	utils.SendResponse(w, http.StatusOK, product)
}
