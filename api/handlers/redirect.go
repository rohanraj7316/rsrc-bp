package handlers

import (
	"dwarf/api/models"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

func Redirect(rw http.ResponseWriter, r *http.Request) (err error) {

	c := r.Context()

	filter := bson.M{"hash": chi.URLParam(r, "hash")}
	data, err := models.Find(c, filter)
	if err != nil {
		return err
	}

	url := data["urlOriginal"]

	http.Redirect(rw, r, url.(string), http.StatusSeeOther)

	return
}
