package product

import "scalvid/domain"

type service struct {
	productRepo ProductRepository
}

func NewService(repo ProductRepository) ProductService {
	return &service{productRepo: repo}
}

// CreateProduct implements Service.
func (svc *service) CreateProduct(product domain.Product) (*domain.Product, error) {
	createdProduct, err := svc.productRepo.Create(product)
	if err != nil {
		return nil, err
	}
	if createdProduct == nil {
		return nil, nil
	}
	return createdProduct, nil
}

// GetProduct implements Service.
func (svc *service) GetProduct(id int) (*domain.Product, error) {
	product, err := svc.productRepo.Get(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}
	return product, nil
}

// GetProducts implements Service.
func (svc *service) GetProducts() ([]*domain.Product, error) {
	products, err := svc.productRepo.Gets()
	if err != nil {
		return nil, err
	}
	if products == nil {
		return nil, nil
	}
	return products, nil
}

// UpdateProduct implements Service.
func (svc *service) UpdateProduct(id int, product domain.Product) (*domain.Product, error) {
	updatedProduct, err := svc.productRepo.Update(id, product)
	if err != nil {
		return nil, err
	}
	if updatedProduct == nil {
		return nil, nil
	}
	return updatedProduct, nil
}

// DeleteProduct implements Service.
func (svc *service) DeleteProduct(id int) error {
	err := svc.productRepo.Delete(id)
	return err
}
