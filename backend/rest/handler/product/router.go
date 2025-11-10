package product

import (
	"net/http"
	"scalvid/rest/middleware"
)

func (h *Handler) RegisterRoute(mux *http.ServeMux, middlewareManager *middleware.Manager) {

	// Product Routes
	mux.Handle(
		"POST /products",
		middlewareManager.With(http.HandlerFunc(h.CreateProductHandler), h.middlewares.AuthenticationJWT),
	)

	mux.Handle(
		"GET /products",
		middlewareManager.With(http.HandlerFunc(h.GetProductsHandler), h.middlewares.AuthenticationJWT),
	)

	mux.Handle(
		"GET /products/{id}",
		middlewareManager.With(http.HandlerFunc(h.GetProductHandler), h.middlewares.AuthenticationJWT),
	)

	mux.Handle(
		"PUT /products/{id}",
		middlewareManager.With(http.HandlerFunc(h.UpdateProductHandler), h.middlewares.AuthenticationJWT),
	)

	mux.Handle(
		"DELETE /products/{id}",
		middlewareManager.With(http.HandlerFunc(h.DeleteProductHandler), h.middlewares.AuthenticationJWT),
	)
}
