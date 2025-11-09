package rest

import (
	"log"
	"net/http"
	"scalvid/config"
	"scalvid/rest/handler/product"
	"scalvid/rest/middleware"
	"strconv"
)

type Server struct {
	productHandler *product.Handler
}

func NewServer(productHandler *product.Handler) *Server {
	return &Server{
		productHandler: productHandler,
	}
}

func (s *Server) StartServer(config config.Config) {
	middlewareManager := middleware.NewManager()
	middlewareManager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := middlewareManager.WrapGlobals(mux)

	// Register Routes
	s.productHandler.RegisterRoute(mux, middlewareManager)

	log.Println("Server is running on port ", config.HttpPort)

	addr := ":" + strconv.Itoa(config.HttpPort)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		log.Println("Error starting server")
		log.Fatal(err)
	}
}
