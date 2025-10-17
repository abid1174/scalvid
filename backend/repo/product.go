package repo

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	ImageUrl    string  `json:"imageUrl"`
}

type ProductRepository interface {
	CreateProduct(product Product) (Product, error)
	GetProduct(id int) (Product, error)
	GetProducts() ([]Product, error)
	UpdateProduct(id int, product Product) (Product, error)
	DeleteProduct(id int) error
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product Product) (Product, error) {
	log.Println("CreateProduct", product)
	rows, err := r.db.NamedQuery(`
		INSERT INTO products (title, price, description, category, image_url)
		VALUES (${Title}, ${Price}, ${Description}, ${Category}, ${ImageUrl})
		RETURNING id
	`, product)

	log.Println(rows)

	if err != nil {
		return Product{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&product)
	}
	return product, err
}

func (r *productRepository) GetProduct(id int) (Product, error) {
	return Product{}, nil
}

func (r *productRepository) GetProducts() ([]Product, error) {
	return []Product{}, nil
}

func (r *productRepository) UpdateProduct(id int, product Product) (Product, error) {
	return Product{}, nil
}

func (r *productRepository) DeleteProduct(id int) error {
	return nil
}
