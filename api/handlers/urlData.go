package handlers

import (
	"dwarf/api/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// CreateURL responsible for creating short url.
func CreateURL(rw http.ResponseWriter, r *http.Request) (err error) {

	c := r.Context()

	decoder := json.NewDecoder(r.Body)
	var rBody = make(map[string]interface{})

	err = decoder.Decode(&rBody)
	if err != nil {
		return err
	}

	data, err := models.CreateURL(c, rBody)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

	return
}

// FindURLByID responsible for fetching data from db.
func FindURLByID(rw http.ResponseWriter, r *http.Request) (err error) {

	c := r.Context()

	id := chi.URLParam(r, "id")

	data, err := models.FindURLByID(c, id)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

	return
}

// UpdateURLByID responsible for updating data into db.
func UpdateURLByID(rw http.ResponseWriter, r *http.Request) (err error) {

	c := r.Context()

	id := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var rBody = make(map[string]interface{})

	err = decoder.Decode(&rBody)
	if err != nil {
		return err
	}

	data, err := models.UpdateURLByID(c, id, rBody)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

	return
}
