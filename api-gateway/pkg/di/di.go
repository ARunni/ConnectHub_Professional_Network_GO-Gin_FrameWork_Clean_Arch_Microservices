package di

import (
	server "connectHub_gateway/pkg/api"
	"connectHub_gateway/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {
	serverHTTP := server.NewServerHTTP()
	return serverHTTP, nil
}
