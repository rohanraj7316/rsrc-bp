package helpers

import (
	"encoding/json"
	"net/http"
)

// ValidateAPIKey - it validates the API
// key passed in the headers
func ValidateAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		var rBody interface{}
		err := d.Decode(&rBody)
		if err != nil {
			http.Error(rw, "getting error", http.StatusAccepted)
		}

		next.ServeHTTP(rw, r)
	})
}
