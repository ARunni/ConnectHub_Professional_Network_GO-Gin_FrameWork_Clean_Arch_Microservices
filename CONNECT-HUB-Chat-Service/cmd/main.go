package main

import (
	"log"

	logging "github.com/ARunni/ConnetHub_chat/Logging"
	"github.com/ARunni/ConnetHub_chat/pkg/config"
	"github.com/ARunni/ConnetHub_chat/pkg/di"
)

func main() {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_chat.log")
	defer logrusLogFile.Close()
	logrusLogger.Info("connectHub_chat main file started")
	logrusLogger.Info("Loading config started")
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		logrusLogger.Error("Failed to load config: ", cfgErr)
		log.Fatal("canot load config: ", cfgErr)
	}
	logrusLogger.Info("Loading config finished")
	logrusLogger.Info("Loading InitializeAPI started")
	server, diErr := di.InitializeAPI(cfg)
	if diErr != nil {
		logrusLogger.Fatal("Cannot start server: ", diErr)
		log.Fatal("cannot start server: ", diErr)
	}
	logrusLogger.Info("Loading InitializeAPI finished")
	logrusLogger.Info("connectHub_chat started")
	server.Start()
}
