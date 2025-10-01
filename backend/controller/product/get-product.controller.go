package controller

import (
	"net/http"
	"scalvid/data"
	"scalvid/utils"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, http.StatusOK, data.Products)
}
