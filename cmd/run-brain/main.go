package main

import (
	"flag"
	"log"

	"leagueapi.com.br/brain/src/entities"

	"leagueapi.com.br/brain/src/infrastructure/server"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	server := server.NewServer().BindURL().BindHeader().Connect()
	entities.NewBrain().StartHandlers().Handle()
	go server.Reciver()
	server.Handler()
}
