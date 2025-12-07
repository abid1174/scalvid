package repo

import (
	"database/sql"

	"scalvid/domain"
	"scalvid/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepository
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepo {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product domain.Product) (*domain.Product, error) {
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

func (r *productRepository) Get(id int) (*domain.Product, error) {
	var product domain.Product

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

func (r *productRepository) Gets() ([]*domain.Product, error) {
	var products []*domain.Product

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

func (r *productRepository) Update(id int, product domain.Product) (*domain.Product, error) {
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

func (r *productRepository) Delete(id int) error {
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
