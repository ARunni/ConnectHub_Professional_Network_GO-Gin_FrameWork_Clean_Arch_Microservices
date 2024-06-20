package service

import (
	"context"
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

func (ns *NotificationServer) UserData(ctx context.Context, Req *pb.UserDataRequest) (*pb.UserDataResponse, error) {
	ns.Logger.Info("UserData at NotificationServer started")
	userId := Req.Id
	ns.Logger.Info("UserData at notificationUsecase started")
	data, err := ns.notificationUsecase.UserData(int(userId))
	if err != nil {
		ns.Logger.Error("error from notificationUsecase", err)
		return nil, err
	}
	ns.Logger.Info("UserData at notificationUsecase finished")
	ns.Logger.Info("UserData at NotificationServer finished")
	return &pb.UserDataResponse{
		Id:       int64(data.UserId),
		Username: data.Username,
	}, nil
}
