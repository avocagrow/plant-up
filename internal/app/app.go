package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/avocagrow/plant-up/internal/db"
	"github.com/avocagrow/plant-up/internal/home"
	"github.com/avocagrow/plant-up/internal/plants"
	"github.com/avocagrow/plant-up/internal/server"
)

var publicRoutes map[string]chi.Router

func init() {
}

type App struct {
	name    string
	db      db.DataStore
	Server  *server.Server
	Routers []chi.Router
	Logger  *slog.Logger
}

type AppOptFunc func(*App) error

func NewApp(opts ...AppOptFunc) (*App, error) {
	a := &App{
		name:   "avocagrow-api-service",
		Logger: slog.Default(),
	}

	db, err := db.Configure()
	a.db = db
	if err != nil {
		slog.Error("unable to configure db connection\n\t%v\n", err)
		os.Exit(1)
	}

	for _, fn := range opts {
		if err = fn(a); err != nil {
			return nil, err
		}
	}

	ps, err := plants.NewPlantService(a.db)
	if err != nil {
		slog.Error("unable to create new plant service\n\t%v\n", err)
		os.Exit(1)
	}

	plantsHandler := plants.NewHandler(ps)

	routes := map[string]chi.Router{
		"/":       home.Routes(),
		"/plants": plants.Routes(plantsHandler),
	}
	r := NewRouter(routes)

	srv, err := server.NewServer(
		a.name,
		server.WithHandler(r),
		server.WithAddr(":8080"),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create new server for app: %w", err)
	}
	slog.Info("created new server with address", slog.Any("server-addr", srv.Addr()))

	a.Server = srv

	return a, nil
}

func WithName(name string) AppOptFunc {
	return func(a *App) error {
		a.name = name
		return nil
	}
}

func WithConnection(ctx context.Context, conn string) AppOptFunc {
	return func(a *App) error {
		return a.db.SetNewConnection(ctx, conn)
	}
}
