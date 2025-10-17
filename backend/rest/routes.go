package rest

import (
	"net/http"
	handler "scalvid/rest/handler/product"
	"scalvid/rest/middleware"
)

func initRoutes(mux *http.ServeMux, middlewareManager *middleware.Manager) {

	// Product Routes
	mux.Handle(
		"POST /products",
		middlewareManager.With(http.HandlerFunc(handler.CreateProductHandler)),
	)

	mux.Handle(
		"GET /products",
		middlewareManager.With(http.HandlerFunc(handler.GetProductsHandler)),
	)

	mux.Handle(
		"GET /products/{id}",
		middlewareManager.With(http.HandlerFunc(handler.GetProductHandler)),
	)

	mux.Handle(
		"PUT /products/{id}",
		middlewareManager.With(http.HandlerFunc(handler.UpdateProductHandler)),
	)

	mux.Handle(
		"DELETE /products/{id}",
		middlewareManager.With(http.HandlerFunc(handler.DeleteProductHandler)),
	)
}
