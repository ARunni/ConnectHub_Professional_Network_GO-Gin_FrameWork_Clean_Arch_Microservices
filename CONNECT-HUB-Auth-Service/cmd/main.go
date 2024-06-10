package main

import (
	logging "github.com/ARunni/ConnetHub_auth/Logging"
	"github.com/ARunni/ConnetHub_auth/pkg/config"
	"github.com/ARunni/ConnetHub_auth/pkg/di"
	"log"
)

func main() {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
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
	logrusLogger.Info("connectHub_auth started")
	server.Start()
}
