package main

import (
	"connectHub_gateway/pkg/config"
	"connectHub_gateway/pkg/di"
	"log"
)

func main() {
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		log.Fatal("canot load config: ", cfgErr)
	}
	server, diErr := di.InitializeAPI(cfg)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	}
	server.Start()
}
