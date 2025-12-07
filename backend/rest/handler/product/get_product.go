package product

import (
	"net/http"
	"scalvid/utils"
)

func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetProducts()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to get products")
		return
	}
	utils.SendResponse(w, http.StatusOK, products)
}
