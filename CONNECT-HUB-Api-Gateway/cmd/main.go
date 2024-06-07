package main

import (
	logging "connectHub_gateway/Logging"
	"connectHub_gateway/pkg/config"
	"connectHub_gateway/pkg/di"
	"log"
)

func main() {

	// logrusLogger, logrusLogFile := logging.InitLogrusLogger("connectHub_gateway.log")
	// defer logrusLogFile.Close()

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		logrusLogger.Error("Failed to load config: ", cfgErr)
		log.Fatal("canot load config: ", cfgErr)
	}
	server, diErr := di.InitializeAPI(cfg)
	if diErr != nil {
		logrusLogger.Fatal("Cannot start server: ", diErr)
		log.Fatal("cannot start server: ", diErr)
	}
	logrusLogger.Info("connectHub_gateway started")
	server.Start()
	

}
