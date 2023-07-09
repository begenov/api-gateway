package server

import (
	"context"
	"net/http"

	"github.com/begenov/api-gateway/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.ServerConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP_Port,
			Handler:        handler,
			ReadTimeout:    cfg.READ_TIMEOUT,
			WriteTimeout:   cfg.WRITE_TIMEOUT,
			MaxHeaderBytes: cfg.MAX_HEADER_BYTES << 20,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
