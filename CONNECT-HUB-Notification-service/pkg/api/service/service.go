package service

import (
	"context"
	"os"

	logging "github.com/ARunni/ConnetHub_Notification/Logging"
	pb "github.com/ARunni/ConnetHub_Notification/pkg/pb/notification"
	interfaces "github.com/ARunni/ConnetHub_Notification/pkg/usecase/interface"
	"github.com/ARunni/ConnetHub_Notification/pkg/utils/models"
	"github.com/sirupsen/logrus"
)

type NotificationServer struct {
	notificationUsecase interfaces.NotificationUseCase
	pb.UnimplementedNotificationServiceServer
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewNotificationServer(usecase interfaces.NotificationUseCase) pb.NotificationServiceServer {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Notification.log")
	return &NotificationServer{
		notificationUsecase: usecase,
		Logger:              logger,
		LogFile:             logFile,
	}
}

func (ad *NotificationServer) GetNotification(ctx context.Context, req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	ad.Logger.Info("GetNotification at NotificationServer started")
	userid := req.UserID

	result, err := ad.notificationUsecase.GetNotification(int(userid), models.Pagination{Limit: int(req.Limit), Offset: int(req.Offset)})
	if err != nil {
		ad.Logger.Error("error from notificationUsecase", err)
		return nil, err
	}
	ad.Logger.Info("GetNotification at notificationUsecase success")
	var final []*pb.Message

	for _, v := range result {
		final = append(final, &pb.Message{
			UserId:   int64(v.UserID),
			Username: v.Username,
			Profile:  v.Profile,
			Message:  v.Message,
			Time:     v.CreatedAt,
		})
	}
	ad.Logger.Info("GetNotification at NotificationServer success")
	return &pb.GetNotificationResponse{
		Notification: final,
	}, nil
}