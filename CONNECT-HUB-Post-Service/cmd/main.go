package main

import (
	"ConnetHub_post/pkg/config"
	"ConnetHub_post/pkg/di"
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
