package routers

import (
	"dwarf/api/handlers"

	"github.com/go-chi/chi"
)

func URLData(r chi.Router) {
	r.Post("/", h(handlers.CreateURL))
	r.Get("/{id}", h(handlers.FindURLByID))
	r.Patch("/{id}", h(handlers.UpdateURLByID))
}
