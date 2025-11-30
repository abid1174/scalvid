package product

import (
	"net/http"
	"scalvid/rest/middleware"
)

func (h *Handler) RegisterRoute(mux *http.ServeMux, middlewareManager *middleware.Manager) {

	// Product Routes
	mux.Handle(
		"POST /products",
		middlewareManager.With(http.HandlerFunc(h.CreateProductHandler)),
	)

	mux.Handle(
		"GET /products",
		middlewareManager.With(http.HandlerFunc(h.GetProductsHandler)),
	)

	mux.Handle(
		"GET /products/{id}",
		middlewareManager.With(http.HandlerFunc(h.GetProductByIdHandler)),
	)

	mux.Handle(
		"PUT /products/{id}",
		middlewareManager.With(http.HandlerFunc(h.UpdateProductHandler)),
	)

	mux.Handle(
		"DELETE /products/{id}",
		middlewareManager.With(http.HandlerFunc(h.DeleteProductHandler)),
	)
}
