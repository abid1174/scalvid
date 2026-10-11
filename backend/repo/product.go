package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

// ErrNotFound is returned when the requested record does not exist.
var ErrNotFound = errors.New("record not found")

type ProductRepository interface {
	CreateProduct(ctx context.Context, product Product) (*Product, error)
	GetProduct(ctx context.Context, id int) (*Product, error)
	GetProducts(ctx context.Context) ([]*Product, error)
	UpdateProduct(ctx context.Context, id int, product Product) (*Product, error)
	DeleteProduct(ctx context.Context, id int) error
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

func (r *productRepository) CreateProduct(ctx context.Context, product Product) (*Product, error) {
	query := `
		INSERT INTO products (title, description, price, image_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	row := r.db.QueryRowContext(ctx, query, product.Title, product.Description, product.Price, product.ImageUrl)

	err := row.Scan(&product.ID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetProduct(ctx context.Context, id int) (*Product, error) {
	var product Product

	query := `
		SELECT id, title, COALESCE(description, '') AS description, price, COALESCE(image_url, '') AS image_url
		FROM products
		WHERE id = $1
	`

	err := r.db.GetContext(ctx, &product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetProducts(ctx context.Context) ([]*Product, error) {
	// Non-nil so an empty table encodes as [] instead of null
	products := []*Product{}

	query := `
		SELECT id, title, COALESCE(description, '') AS description, price, COALESCE(image_url, '') AS image_url
		FROM products
		ORDER BY id
	`

	err := r.db.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, id int, product Product) (*Product, error) {
	var updated Product

	query := `
		UPDATE products
		SET title = $1, description = $2, price = $3, image_url = $4, updated_at = NOW()
		WHERE id = $5
		RETURNING id, title, COALESCE(description, '') AS description, price, COALESCE(image_url, '') AS image_url
	`

	err := r.db.GetContext(ctx, &updated, query, product.Title, product.Description, product.Price, product.ImageUrl, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &updated, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, id int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
