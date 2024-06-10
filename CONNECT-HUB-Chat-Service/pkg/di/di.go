package di

import (
	server "github.com/ARunni/ConnetHub_chat/pkg/api/server"
	"github.com/ARunni/ConnetHub_chat/pkg/api/service"
	"github.com/ARunni/ConnetHub_chat/pkg/client"
	"github.com/ARunni/ConnetHub_chat/pkg/config"
	"github.com/ARunni/ConnetHub_chat/pkg/db"
	"github.com/ARunni/ConnetHub_chat/pkg/repository"
	"github.com/ARunni/ConnetHub_chat/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	database, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	chatRepository := repository.NewChatRepository(database)
	authClient := client.NewAuthClient(&cfg)

	chatUseCase := usecase.NewChatUseCase(chatRepository, authClient.Client)

	serviceServer := service.NewChatServer(chatUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, serviceServer)
	if err != nil {
		return nil, err
	}

	go chatUseCase.MessageConsumer()
	return grpcServer, nil
}
