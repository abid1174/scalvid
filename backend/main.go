package main

import (
	"log"
	"net/http"
	controller "scalvid/controller/product"
	"scalvid/middleware"
)

func main() {
	// Initialize routes mux
	mux := http.NewServeMux()

	// Initialize middleware manager
	middlewareManager := middleware.NewManager()
	middlewares := middlewareManager.With(middleware.Cors, middleware.Logger, middleware.Test)

	// Register routes
	mux.Handle("GET /", middlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})))

	mux.Handle("GET /products", middlewares(http.HandlerFunc(controller.GetProductsHandler)))
	mux.Handle("POST /create-product", middlewares(http.HandlerFunc(controller.CreateProductHandler)))

	log.Println("Server is running on port 8000")

	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}
}
