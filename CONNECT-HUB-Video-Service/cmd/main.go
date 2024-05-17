package main

import (
	"ConnetHub_video/pkg/config"
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
