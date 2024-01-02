package plants

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(h *Handler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.listPlants)
	r.Post("/", h.createPlant)
	r.Patch("/{plantId}", h.updatePlant)
	return r
}

type Handler struct {
	p Planter
}

// NewHandler returns an instance of the Product handler
func NewHandler(p Planter) *Handler {
	return &Handler{p: p}
}

func (h *Handler) listPlants(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Plantheads!"))
}

func (h *Handler) createPlant(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handler) updatePlant(w http.ResponseWriter, r *http.Request) {
}
