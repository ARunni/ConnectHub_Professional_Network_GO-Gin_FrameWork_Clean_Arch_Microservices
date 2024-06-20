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
			Id:       int64(v.ID),
			Message:  v.Message,
			Time:     v.CreatedAt,
			PostId:   int64(v.PostID),
		})
	}
	ad.Logger.Info("GetNotification at NotificationServer success")
	return &pb.GetNotificationResponse{
		Notification: final,
	}, nil
}

func (ad *NotificationServer) ReadNotification(ctx context.Context, req *pb.ReadNotificationRequest) (*pb.ReadNotificationResponse, error) {
	ad.Logger.Info("ReadNotification at NotificationServer started")
	userid := req.UserId
	id := req.Id

	result, err := ad.notificationUsecase.ReadNotification(int(id), int(userid))
	if err != nil {
		ad.Logger.Error("error from notificationUsecase", err)
		return nil, err
	}
	ad.Logger.Info("ReadNotification at notificationUsecase success")

	ad.Logger.Info("GetNotification at NotificationServer success")
	return &pb.ReadNotificationResponse{
		Success: result,
	}, nil
}

func (ad *NotificationServer) MarkAllAsRead(ctx context.Context, req *pb.MarkAllAsReadRequest) (*pb.MarkAllAsReadResponse, error) {
	ad.Logger.Info("MarkAllAsRead at NotificationServer started")
	userid := req.UserId

	result, err := ad.notificationUsecase.MarkAllAsRead(int(userid))
	if err != nil {
		ad.Logger.Error("error from notificationUsecase", err)
		return nil, err
	}
	ad.Logger.Info("MarkAllAsRead at notificationUsecase success")

	ad.Logger.Info("MarkAllAsRead at NotificationServer success")
	return &pb.MarkAllAsReadResponse{
		Success: result,
	}, nil
}
