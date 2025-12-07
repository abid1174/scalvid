package cmd

import (
	"scalvid/config"
	"scalvid/infra/db"
	"scalvid/product"
	"scalvid/repo"
	"scalvid/rest"
	prdHandler "scalvid/rest/handler/product"
	"scalvid/rest/middleware"
)

func Serve() {
	config := config.GetConfig()

	dbConnection := db.NewConnection(config.DB)
	productRepository := repo.NewProductRepository(dbConnection)

	productService := product.NewService(productRepository)

	middlewares := middleware.NewMiddlewares(config)

	productHandler := prdHandler.NewHandler(middlewares, productService)

	server := rest.NewServer(config, productHandler)
	server.StartServer()
}
