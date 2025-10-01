package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	ImageUrl    string  `json:"imageUrl"`
}

var products []Product

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hanldeCORS(w)
	handlePreflight(w, r)

	w.Write([]byte("Hello, World!"))
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	hanldeCORS(w)
	handlePreflight(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	hanldeCORS(w)
	handlePreflight(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	product.ID = len(products) + 1
	log.Println("Product created: ", product)
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusCreated)
}

func hanldeCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

}

func handlePreflight(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/products", getProductsHandler)
	mux.HandleFunc("/create-product", createProductHandler)

	log.Println("Server is running on port 8000")

	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Product 1",
		Price:       100,
		Description: "Product 1 description",
		Category:    "Category 1",
		ImageUrl:    "https://via.placeholder.com/150",
	}

	product2 := Product{
		ID:          2,
		Title:       "Product 2",
		Price:       200,
		Description: "Product 2 description",
		Category:    "Category 2",
		ImageUrl:    "https://via.placeholder.com/150",
	}

	products = append(products, product1, product2)
}
