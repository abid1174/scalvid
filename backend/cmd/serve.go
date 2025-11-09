package cmd

import (
	"scalvid/config"
	"scalvid/rest"
	"scalvid/rest/handler/product"
)

func Serve() {
	config := config.GetConfig()

	productHandler := product.NewHandler()

	server := rest.NewServer(productHandler)
	server.StartServer(config)
}
