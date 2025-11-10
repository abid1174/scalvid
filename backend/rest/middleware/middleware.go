package middleware

import "scalvid/config"

type Middlewares struct {
	config *config.Config
}

func NewMiddlewares(cfg *config.Config) *Middlewares {
	return &Middlewares{
		config: cfg,
	}
}
