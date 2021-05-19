package routers

import (
	"dwarf/api/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// RouterHandler - generic router handler
func RouterHandler() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/", Redirect)
	r.Route("/url-shortener", URLData)
	return r
}

func h(fn handlers.Handler) http.HandlerFunc {
	return handlers.Handler(fn).ErrorHandler
}
