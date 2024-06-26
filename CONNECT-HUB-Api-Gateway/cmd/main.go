package main

import (
	"log"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	"github.com/ARunni/connectHub_gateway/pkg/config"
	"github.com/ARunni/connectHub_gateway/pkg/di"

	_ "github.com/ARunni/connectHub_gateway/cmd/docs"

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
// @description For accessing the API, you need to include "Jobseeker","Admin" or "Recruiter" before your token in the Authorization header. Example: "Authorization: Admin <your_token>"
// @host localhost:7000
// @BasePath /
// @query.collection.format multi
func main() {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	logrusLogger.Info("connectHub_gateway main file started")
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
	logrusLogger.Info("connectHub_gateway started")

	server.Start()

}
