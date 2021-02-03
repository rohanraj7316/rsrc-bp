package routers

import (
	"dwarf/api/handlers"

	"github.com/go-chi/chi"
)

// Health - health route
func Health(r chi.Router) {
	r.Get("/health", h(handlers.GetHealth))
}
