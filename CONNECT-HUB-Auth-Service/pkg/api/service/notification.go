package service

import (
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	pb "github.com/ARunni/ConnetHub_auth/pkg/pb/notification/auth"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"
	"github.com/sirupsen/logrus"
)

type NotificationServer struct {
	notificationUsecase interfaces.NotificationUsecase
	pb.UnimplementedNotificationAuthServiceServer
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewNotificationServer(useCase interfaces.NotificationUsecase) pb.NotificationAuthServiceServer {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &NotificationServer{
		notificationUsecase: useCase,
		Logger:              logger,
		LogFile:             logFile,
	}
}
