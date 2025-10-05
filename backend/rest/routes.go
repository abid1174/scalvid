package rest

import (
	"net/http"
	controller "scalvid/rest/controller/product"
	"scalvid/rest/middleware"
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
