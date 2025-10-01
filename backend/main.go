package main

import (
	"log"
	"net/http"
	controller "scalvid/controller/product"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})))

	mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(controller.GetProductsHandler)))
	mux.Handle("POST /create-product", corsMiddleware(http.HandlerFunc(controller.CreateProductHandler)))

	log.Println("Server is running on port 8000")

	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}
}

// ========= Middleware ================

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
