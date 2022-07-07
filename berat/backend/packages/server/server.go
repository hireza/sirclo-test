package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hireza/sirclo-test/berat/packages/config"
	"github.com/rs/zerolog/log"
)

type Server struct {
	http   *http.Server
	Router *mux.Router
}

func NewServer(cfg *config.Config) *Server {
	r := mux.NewRouter()

	return &Server{
		http: &http.Server{
			Addr: cfg.Server.Port,
		},
		Router: r,
	}
}

func (s *Server) ListenAndServe() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	errc := make(chan error)
	go func() {
		log.Printf("HTTP Server listen on %s", s.http.Addr)
		errc <- s.http.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case err := <-errc:
		log.Error().Err(err).Msg(fmt.Sprintf("error when listen on %s\n", s.http.Addr))
		return err
	case <-quit:
		log.Printf("Shutting down the server")
		return s.http.Shutdown(ctx)
	}
}

func (s *Server) RegisterRouter(handler http.Handler) {
	s.http.Handler = handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"Content-Type"}))(handler)
}
