package di

import (
	server "github.com/ARunni/ConnetHub_Notification/pkg/api/server"
	"github.com/ARunni/ConnetHub_Notification/pkg/api/service"
	client "github.com/ARunni/ConnetHub_Notification/pkg/client/auth"
	"github.com/ARunni/ConnetHub_Notification/pkg/config"
	"github.com/ARunni/ConnetHub_Notification/pkg/db"
	"github.com/ARunni/ConnetHub_Notification/pkg/repository"
	"github.com/ARunni/ConnetHub_Notification/pkg/usecase"
)

func InitializeApi(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	notiRepository := repository.NewNotificationRepository(gormDB)
	noticlient := client.NewAuthClient(&cfg)
	notiUseCase := usecase.NewNotificationUsecase(notiRepository, noticlient)
	notiServiceServer := service.NewNotificationServer(notiUseCase)
	grpcserver, err := server.NewGRPCServer(cfg, notiServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	go notiUseCase.ConsumeNotification()
	return grpcserver, nil
}
