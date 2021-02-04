package main

import (
	"dwarf/api/routers"
	"log"
	"net/http"

	"dwarf/api/helpers"
)

//Connection mongoDB with helper class
var collection = helpers.ConnectDB()

func main() {

	// handling server error here.
	defer func() {
		if err := recover(); err != nil {
			log.SetPrefix("[Error]")
			log.Panicln(err)
		}
	}()

	mux := routers.RouterHandler()

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	panic(server.ListenAndServe())
}
