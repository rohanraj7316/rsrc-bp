package main

import (
	"context"
	"dwarf/api/routers"
	"dwarf/api/services"
	"log"
	"math/rand"
	"net/http"
	"time"

	"dwarf/configs"
)

var serverConfig configs.ServerConfigStruct

func init() {
	rand.Seed(time.Now().UnixNano())

	err := configs.Initialize(&serverConfig)
	if err != nil {
		log.Panicln(err)
	}
}

func main() {

	// handling server error here.
	defer func() {
		if err := recover(); err != nil {
			log.SetPrefix("[Error]")
			log.Panicln(err)
		}
	}()

	c := context.Background()

	services.MongoConnect(c)

	mux := routers.RouterHandler()

	server := http.Server{
		Addr:    ":" + serverConfig.Port,
		Handler: mux,
	}

	log.Println("server started successfully")
	panic(server.ListenAndServe())
}
