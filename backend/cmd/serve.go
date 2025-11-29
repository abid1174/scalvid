package cmd

import (
	"scalvid/config"
	"scalvid/infra/db"
	"scalvid/repo"
	"scalvid/rest"
	"scalvid/rest/handler/product"
	"scalvid/rest/middleware"
)

func Serve() {
	config := config.GetConfig()

	dbConnection := db.NewConnection()
	productRepository := repo.NewProductRepository(dbConnection)

	middlewares := middleware.NewMiddlewares(config)

	productHandler := product.NewHandler(middlewares, productRepository)

	server := rest.NewServer(config, productHandler)
	server.StartServer()
}
