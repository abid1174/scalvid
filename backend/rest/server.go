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
	config         *config.Config
	productHandler *product.Handler
}

func NewServer(config *config.Config, productHandler *product.Handler) *Server {
	return &Server{
		config:         config,
		productHandler: productHandler,
	}
}

func (s *Server) StartServer() {
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

	log.Println("Server is running on port ", s.config.HttpPort)

	addr := ":" + strconv.Itoa(s.config.HttpPort)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		log.Println("Error starting server")
		log.Fatal(err)
	}
}
