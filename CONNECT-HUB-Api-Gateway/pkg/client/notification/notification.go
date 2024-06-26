package client

import (
	"context"
	"os"

	logging "github.com/ARunni/connectHub_gateway/Logging"

	client "github.com/ARunni/connectHub_gateway/pkg/client/notification/interfaces"
	"github.com/ARunni/connectHub_gateway/pkg/config"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"

	"fmt"

	Pb "github.com/ARunni/connectHub_gateway/pkg/pb/notification"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type notificationClient struct {
	Client  Pb.NotificationServiceClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewNotificationClient(cfg config.Config) client.NotificationClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	grpcConnection, err := grpc.Dial(cfg.Connect_Hub_Notification, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth", err)
	}
	grpcClient := Pb.NewNotificationServiceClient(grpcConnection)
	return &notificationClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}

}

func (nc *notificationClient) GetNotification(userid int, mod models.NotificationPagination) ([]models.NotificationResponse, error) {
	nc.Logger.Info("GetNotification at notificationClient started")
	nc.Logger.Info("GetNotification at client started")

	data, err := nc.Client.GetNotification(context.Background(), &Pb.GetNotificationRequest{
		UserID: int64(userid),
		Limit:  int64(mod.Limit),
		Offset: int64(mod.Offset),
	})
	if err != nil {
		nc.Logger.Error("error from client ", err)
		return []models.NotificationResponse{}, err
	}
	nc.Logger.Info("GetNotification at client finished")
	var response []models.NotificationResponse
	for _, v := range data.Notification {
		notificationResponse := models.NotificationResponse{
			ID:        int(v.Id),
			UserID:    int(v.UserId),
			Username:  v.Username,
			PostID:    int(v.PostId),
			Message:   v.Message,
			CreatedAt: v.Time,
		}
		response = append(response, notificationResponse)
	}
	return response, nil

}

func (nc *notificationClient) ReadNotification(id, user_id int) (bool, error) {
	nc.Logger.Info("ReadNotification at notificationClient started")
	nc.Logger.Info("ReadNotification at client started")
	ok, err := nc.Client.ReadNotification(context.Background(), &Pb.ReadNotificationRequest{
		UserId: int64(user_id),
		Id:     int64(id),
	})
	if err != nil {
		nc.Logger.Error("error from GRPC call", err)
		return false, err
	}
	nc.Logger.Info("ReadNotification at notificationClient finished")
	nc.Logger.Info("ReadNotification at client finished")
	return ok.Success, nil
}

func (nc *notificationClient) MarkAllAsRead(userId int) (bool, error) {
	nc.Logger.Info("MarkAllAsRead at notificationClient started")
	nc.Logger.Info("MarkAllAsRead at client started")
	ok, err := nc.Client.MarkAllAsRead(context.Background(), &Pb.MarkAllAsReadRequest{
		UserId: int64(userId),
	})
	if err != nil {
		nc.Logger.Error("error from GRPC call", err)
		return false, err
	}
	nc.Logger.Info("MarkAllAsRead at notificationClient finished")
	nc.Logger.Info("MarkAllAsRead at client finished")
	return ok.Success, nil
}

func (nc *notificationClient) GetAllNotifications(userId int) ([]models.AllNotificationResponse, error) {
	nc.Logger.Info("GetAllNotifications at notificationClient started")
	nc.Logger.Info("GetAllNotifications at client started")
	data, err := nc.Client.GetAllNotifications(context.Background(), &Pb.GetAllNotificationsRequest{
		UserId: int64(userId),
	})
	if err != nil {
		nc.Logger.Error("error from GRPC call", err)
		return nil, err
	}
	nc.Logger.Info("GetAllNotifications at notificationClient finished")
	var response []models.AllNotificationResponse
	for _, v := range data.Notification {
		notificationResponse := models.AllNotificationResponse{
			ID:        int(v.Id),
			UserID:    int(v.UserId),
			Username:  v.Username,
			PostID:    int(v.PostId),
			Message:   v.Message,
			CreatedAt: v.Time,
			Read:      v.Read,
		}
		response = append(response, notificationResponse)
	}
	nc.Logger.Info("GetAllNotifications at client finished")
	return response, nil
}
