package routers

import (
	"dwarf/api/handlers"

	"github.com/go-chi/chi"
)

func URLData(r chi.Router) {
	r.Get("/{id}", h(handlers.FindURLByID))
	r.Get("/health", h(handlers.GetHealth))

	r.Post("/", h(handlers.CreateURL))

	r.Patch("/{id}", h(handlers.UpdateURLByID))
}
