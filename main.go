package main

import (
	"flag"
	"log"

	"github.com/mathiasdonoso/fakery/internal/fakery"
)

func main() {
	port := flag.String("p", "8000", "Fakery server's port")
	configPath := flag.CommandLine.String("d", "config.json", "Fakery server's configuration")
	flag.Parse()

	config, err := fakery.CreateNewFakeryServerConfig(*configPath)
	if err != nil {
		log.Fatal("Error trying to use the config file", err)
		return
	}

	server := fakery.CreateNewServer(*port, config)
	server.Start()
}
