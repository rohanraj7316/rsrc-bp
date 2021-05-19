package routers

import (
	"dwarf/api/handlers"
	"dwarf/api/helpers"

	"github.com/go-chi/chi"
)

func URLData(r chi.Router) {

	r.Use(helpers.ValidateAPIKey)

	r.Get("/{id}", h(handlers.FindURLByID))
	r.Get("/health", h(handlers.GetHealth))

	r.Post("/", h(handlers.CreateURL))

	r.Patch("/{id}", h(handlers.UpdateURLByID))
}
