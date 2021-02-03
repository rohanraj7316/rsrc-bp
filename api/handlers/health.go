package handlers

import "net/http"

// GetHealth - generic health check
func GetHealth(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	return nil
}
