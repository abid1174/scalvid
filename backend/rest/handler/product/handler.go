package product

import (
	"scalvid/product"
	"scalvid/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	service     product.ProductService
}

func NewHandler(m *middleware.Middlewares, service product.ProductService) *Handler {
	return &Handler{
		middlewares: m,
		service:     service,
	}
}
