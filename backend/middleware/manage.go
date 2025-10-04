package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

// Handle Route Level Middleware
func (m *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}

// Handle Global Middlewares
func (m *Manager) WrapGlobals(handler http.Handler) http.Handler {
	h := handler

	for _, middleware := range m.globalMiddlewares {
		h = middleware(h)
	}

	return h
}
