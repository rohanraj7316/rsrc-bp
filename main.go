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

func init() {
	rand.Seed(time.Now().UnixNano())
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

	serverConfig := configs.ServerConfig{}
	err := configs.Initialize(&serverConfig)
	if err != nil {

	}

	mux := routers.RouterHandler()

	server := http.Server{
		Addr:    ":" + serverConfig.Port,
		Handler: mux,
	}

	log.Println("server started successfully")
	panic(server.ListenAndServe())
}
