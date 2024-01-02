package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Router is a wrapper around chi.Mux
type Router struct {
	*chi.Mux
}

// NewRouter creates a new Chi router with middleware
func NewRouter(routes map[string]chi.Router) *Router {
	r := chi.NewRouter()
	r.Use(middleware.CleanPath)
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json", "application/text"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	for path, fn := range routes {
		r.Mount(path, fn)
	}

	return &Router{r}
}
