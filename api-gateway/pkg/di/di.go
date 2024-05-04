package di

import (
	server "connectHub_gateway/pkg/api"
	"connectHub_gateway/pkg/api/handler"
	"connectHub_gateway/pkg/client"
	"connectHub_gateway/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	serverHTTP := server.NewServerHTTP(adminHandler)
	return serverHTTP, nil
}
