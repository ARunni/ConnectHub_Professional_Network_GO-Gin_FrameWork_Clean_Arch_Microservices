package main

import (
	"ConnetHub_chat/pkg/config"
	"ConnetHub_chat/pkg/di"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("cannot start server:", err)
	} else {
		server.Start()
	}
}
