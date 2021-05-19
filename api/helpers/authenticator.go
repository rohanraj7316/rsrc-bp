package helpers

import (
	"context"
	"dwarf/configs"
	"log"
	"net/http"
)

var authConfig configs.AuthenticatorStruct

func init() {
	err := configs.Initialize(&authConfig)
	if err != nil {
		// TODO: log the error
	}

	if authConfig.AAuthURL == "" {
		log.Println("[Error]")
	}
}

// postAAuthRequest send aauth request
func postAAuthRequest(c context.Context, headers http.Header) (interface{}, error) {

	// extract aauth key from the headers
	apiKey := headers.Get(authConfig.AAuthKey)
	if apiKey == "" {
		// TODO: send error of auth token missing
	}

	// client := http.Client{
	// 	Timeout: time.Duration(),
	// }

	return r, nil
}

// ValidateAPIKey validates API key passed in the headers
func ValidateAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		c := r.Context()

		byPassAAuth := r.URL.Query().Get("byPassAAuth")
		if byPassAAuth == "true" || authConfig.IsAAuthDisabled {
			next.ServeHTTP(rw, r)
		}

		resp, err := postAAuthRequest(c, r.Header)
		if err != nil {

		}
	})
}

// ValidateToken
func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)
	})
}
