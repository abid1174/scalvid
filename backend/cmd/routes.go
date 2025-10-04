package cmd

import (
	"net/http"
	controller "scalvid/controller/product"
	"scalvid/middleware"
)

func initRoutes(mux *http.ServeMux, middlewareManager *middleware.Manager) {

	mux.Handle(
		"POST /create-product",
		middlewareManager.With(http.HandlerFunc(controller.CreateProductHandler), middleware.Test),
	)

	mux.Handle(
		"GET /products",
		middlewareManager.With(http.HandlerFunc(controller.GetProductsHandler)),
	)
}
