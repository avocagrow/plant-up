package home

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", homepage)
	return r
}

// homepage returns the html template for the main page found at "/"
func homepage(w http.ResponseWriter, r *http.Request) {
	slog.InfoContext(r.Context(), "home page was hit!")
	w.Write([]byte("Welcome to the plant-up API!"))
	w.WriteHeader(http.StatusOK)
}
