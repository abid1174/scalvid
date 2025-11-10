package product

import "scalvid/rest/middleware"

type Handler struct {
	middlewares *middleware.Middlewares
}

func NewHandler(m *middleware.Middlewares) *Handler {
	return &Handler{
		middlewares: m,
	}
}
