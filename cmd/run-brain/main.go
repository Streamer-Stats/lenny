package main

import (
	"flag"
	"log"

	"leagueapi.com.br/brain/src/infrastructure/server"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	server := server.NewServer().BindURL().Connect()
	go server.Reciver()
	server.Handler()
}
