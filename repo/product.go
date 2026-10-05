package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	CreateProduct(product Product) (*Product, error)
	GetProduct(id int) (*Product, error)
	GetProducts() ([]*Product, error)
	UpdateProduct(id int, product Product) (*Product, error)
	DeleteProduct(id int) error
}

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
	ImageUrl    string  `json:"imageUrl" db:"image_url"`
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product Product) (*Product, error) {
	query := `
		INSERT INTO products (title, description, price, image_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	row := r.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImageUrl)

	err := row.Scan(&product.ID)
	if err != nil {
		return nil, err
	}

	return &product, err
}

func (r *productRepository) GetProduct(id int) (*Product, error) {
	var product Product

	query := `
		SELECT id, title, description, price, image_url
		FROM products
		WHERE id = $1
	`

	err := r.db.Get(&product, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetProducts() ([]*Product, error) {
	var products []*Product

	query := `
		SELECT id, title, description, price, image_url
		FROM products
	`

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) UpdateProduct(id int, product Product) (*Product, error) {
	query := `
		UPDATE products
		SET title = $1, description = $2, price = $3, image_url = $4
		WHERE id = $5
	`

	row := r.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImageUrl, id)
	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) DeleteProduct(id int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
