package product

import (
	"scalvid/domain"
)

type ProductService interface {
	CreateProduct(product domain.Product) (*domain.Product, error)
	GetProduct(id int) (*domain.Product, error)
	GetProducts() ([]*domain.Product, error)
	UpdateProduct(id int, product domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
}

type ProductRepository interface {
	Create(product domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	Gets() ([]*domain.Product, error)
	Update(id int, product domain.Product) (*domain.Product, error)
	Delete(id int) error
}
