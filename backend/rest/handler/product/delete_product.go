package product

import (
	"net/http"
	"scalvid/utils"
	"strconv"
)

func (h *Handler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.productRepo.DeleteProduct(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}
}
