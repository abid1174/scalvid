package product

import (
	"scalvid/repo"
	"scalvid/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repo.ProductRepository
}

func NewHandler(m *middleware.Middlewares, pr repo.ProductRepository) *Handler {
	return &Handler{
		middlewares: m,
		productRepo: pr,
	}
}
