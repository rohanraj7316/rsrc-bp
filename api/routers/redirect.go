package routers

import (
	"dwarf/api/handlers"

	"github.com/go-chi/chi"
)

func Redirect(r chi.Router) {
	r.Get("/{hash}", h(handlers.Redirect))
}
