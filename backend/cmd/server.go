package cmd

import (
	"log"
	"net/http"
	"scalvid/middleware"
)

func Serve() {
	middlewareManager := middleware.NewManager()
	middlewareManager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := middlewareManager.WrapGlobals(mux)

	initRoutes(mux, middlewareManager)

	log.Println("Server is running on port 8000")

	err := http.ListenAndServe(":8000", wrappedMux)

	if err != nil {
		log.Println("Error starting server")
		log.Fatal(err)
	}
}
