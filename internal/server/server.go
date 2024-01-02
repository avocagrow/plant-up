package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

// ServerOptFunc is a function that can be used to customize a serverer
type ServerOptFunc func(*Server) error

// Server wraps an http.Server
type Server struct {
	name   string
	server *http.Server
}

// WithAddr sets the TCP address to listen on for an http server.
// If one is not provided, http.Server will default to :80.
func WithAddr(addr string) ServerOptFunc {
	return func(s *Server) error {
		s.server.Addr = addr
		return nil
	}
}

// WithHandler sets the http handler for the server.
func WithHandler(h http.Handler) ServerOptFunc {
	return func(s *Server) error {
		s.server.Handler = h
		return nil
	}
}

// NewServer configures and returns a new Server instance. Defaults are
// used when no ServerOptFuncs are provided.
func NewServer(name string, opts ...ServerOptFunc) (*Server, error) {
	srv := &Server{
		name:   name,
		server: &http.Server{},
	}

	for _, fn := range opts {
		if fn == nil {
			continue
		}
		if err := fn(srv); err != nil {
			return nil, fmt.Errorf("unable to create new http.Server due to unexpected error: %w", err)
		}
	}
	return srv, nil
}

func (s *Server) ListenAndServe() {
	slog.Info("listening at:", slog.Any("server-addr", s.server.Addr))
	s.server.ListenAndServe()
}

func (s *Server) Addr() string {
	return s.server.Addr
}
