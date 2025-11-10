package cmd

import (
	"scalvid/config"
	"scalvid/rest"
	"scalvid/rest/handler/product"
	"scalvid/rest/middleware"
)

func Serve() {
	config := config.GetConfig()

	middlewares := middleware.NewMiddlewares(config)

	productHandler := product.NewHandler(middlewares)

	server := rest.NewServer(config, productHandler)
	server.StartServer()
}
