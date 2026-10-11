package product

import (
	"net/http"
	"scalvid/repo"
	"scalvid/utils"
)

func (h *Handler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product repo.Product

	err := utils.DecodeJSON(w, r, &product)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validateProduct(product); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdProduct, err := h.productRepo.CreateProduct(r.Context(), repo.Product{
		Title:       product.Title,
		Price:       product.Price,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}
	utils.SendResponse(w, http.StatusCreated, createdProduct)
}
