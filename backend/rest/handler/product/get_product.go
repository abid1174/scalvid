package product

import (
	"net/http"
	"scalvid/database"
	"scalvid/utils"
)

func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := database.GetProducts()
	utils.SendResponse(w, http.StatusOK, products)
}
