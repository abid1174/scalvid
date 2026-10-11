package rest

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"scalvid/config"
	"scalvid/rest/handler/product"
	"scalvid/rest/middleware"
	"strconv"
	"syscall"
	"time"
)

// Keep below stop_grace_period in docker-compose.yml (20s), or Docker
// SIGKILLs the process before in-flight requests finish.
const shutdownTimeout = 15 * time.Second

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

// StartServer blocks until the server stops. On SIGINT/SIGTERM it stops
// accepting new connections and waits for in-flight requests to finish.
func (s *Server) StartServer() {
	middlewareManager := middleware.NewManager()
	middlewareManager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := middlewareManager.WrapGlobals(mux)

	// Liveness probe
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Register Routes
	s.productHandler.RegisterRoute(mux, middlewareManager)

	server := &http.Server{
		Addr:              ":" + strconv.Itoa(s.config.HttpPort),
		Handler:           wrappedMux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	serverErr := make(chan error, 1)
	go func() {
		log.Println("Server is running on port", s.config.HttpPort)
		serverErr <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("Error starting server")
			log.Fatal(err)
		}
	case <-ctx.Done():
		// Restore default signal handling so a second Ctrl+C force-quits
		stop()
		log.Println("Shutting down server... (press Ctrl+C again to force)")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Println("Graceful shutdown failed:", err)
		}
		log.Println("Server stopped")
	}
}
