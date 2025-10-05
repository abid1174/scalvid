package rest

import (
	"log"
	"net/http"
	"scalvid/config"
	"scalvid/rest/middleware"
	"strconv"
)

func StartServer(config config.Config) {
	middlewareManager := middleware.NewManager()
	middlewareManager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := middlewareManager.WrapGlobals(mux)

	initRoutes(mux, middlewareManager)

	log.Println("Server is running on port ", config.HttpPort)

	addr := ":" + strconv.Itoa(config.HttpPort)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		log.Println("Error starting server")
		log.Fatal(err)
	}
}
