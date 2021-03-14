package handlers

import (
	"dwarf/api/models"
	"encoding/json"
	"net/http"
)

// CreateURL responsible for creating short url.
func CreateURL(rw http.ResponseWriter, r *http.Request) (err error) {

	decoder := json.NewDecoder(r.Body)
	var rBody = make(map[string]interface{})

	err = decoder.Decode(&rBody)
	if err != nil {
		return err
	}

	data, err := models.CreateURL(rBody)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

	return
}
