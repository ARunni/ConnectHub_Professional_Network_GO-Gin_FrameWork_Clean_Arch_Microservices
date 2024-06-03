package di

import (
	server "ConnetHub_chat/pkg/api/server"
	"ConnetHub_chat/pkg/api/service"
	"ConnetHub_chat/pkg/client"
	"ConnetHub_chat/pkg/config"
	"ConnetHub_chat/pkg/db"
	"ConnetHub_chat/pkg/repository"
	"ConnetHub_chat/pkg/usecase"
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
