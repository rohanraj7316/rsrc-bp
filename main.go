package main

import (
	"dwarf/api/routers"
	"log"
	"net/http"

	"dwarf/configs"
)

func main() {

	// handling server error here.
	defer func() {
		if err := recover(); err != nil {
			log.SetPrefix("[Error]")
			log.Panicln(err)
		}
	}()

	serverConfig := configs.ServerConfig{}
	configs.Initialize(&serverConfig)

	mux := routers.RouterHandler()

	server := http.Server{
		Addr:    ":" + serverConfig.PORT,
		Handler: mux,
	}

	log.Println("server started")
	panic(server.ListenAndServe())
}
