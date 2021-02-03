package handlers

import (
	"log"
	"net/http"
)

// Handler - generic struct for handler.
type Handler func(w http.ResponseWriter, r *http.Request) error

// ErrorHandler - handling errors for handlers.
func (fn Handler) ErrorHandler(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		log.SetPrefix("[Error Handler]")
		log.Println(err)
	}
}
