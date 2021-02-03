package routers

import (
	"dwarf/api/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

// RouterHandler - generic router handler
func RouterHandler() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/", Health)
	return r
}

func h(fn handlers.Handler) http.HandlerFunc {
	return handlers.Handler(fn).ErrorHandler
}
