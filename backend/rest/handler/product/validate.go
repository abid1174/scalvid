package product

import (
	"errors"
	"scalvid/repo"
	"strings"
)

func validateProduct(product repo.Product) error {
	if strings.TrimSpace(product.Title) == "" {
		return errors.New("Title is required")
	}

	if product.Price <= 0 {
		return errors.New("Price must be greater than 0")
	}

	return nil
}
