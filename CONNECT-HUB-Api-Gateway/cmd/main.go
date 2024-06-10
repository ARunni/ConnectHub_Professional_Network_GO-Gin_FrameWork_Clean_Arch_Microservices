package main

import (
	logging "connectHub_gateway/Logging"
	"connectHub_gateway/pkg/config"
	"connectHub_gateway/pkg/di"
	"log"

	_ "connectHub_gateway/cmd/docs"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"

)


// @title Go + Gin professional networking platform API Connect Hub
// @version 1.0.0
// @description ConnectHub is a professional networking platform.
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:7000
// @BasePath /
// @query.collection.format multi
func main() {

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
