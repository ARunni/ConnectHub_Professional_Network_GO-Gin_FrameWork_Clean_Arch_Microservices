package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	logging "github.com/ARunni/ConnetHub_Notification/Logging"
	authface "github.com/ARunni/ConnetHub_Notification/pkg/client/auth/interface"
	"github.com/sirupsen/logrus"

	interfaces "github.com/ARunni/ConnetHub_Notification/pkg/repository/interface"
	services "github.com/ARunni/ConnetHub_Notification/pkg/usecase/interface"

	"github.com/ARunni/ConnetHub_Notification/pkg/config"
	"github.com/ARunni/ConnetHub_Notification/pkg/utils/models"
	"github.com/IBM/sarama"
)

type notificationUsecase struct {
	notiRepository interfaces.NotificationRepository
	authclient     authface.Newauthclient
	Logger         *logrus.Logger
	LogFile        *os.File
}

func NewNotificationUsecase(repository interfaces.NotificationRepository, authface authface.Newauthclient) services.NotificationUseCase {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Notification.log")
	return &notificationUsecase{
		notiRepository: repository,
		authclient:     authface,
		Logger:         logger,
		LogFile:        logFile,
	}
}

func (c *notificationUsecase) ConsumeNotification() {

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("error in load config")
	}
	fmt.Println("config at noti kafka", cfg)

	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{cfg.KafkaBrokers}, configs)
	if err != nil {
		fmt.Println("error creating kafka consumer", err)

	}
	fmt.Println("consumer ", consumer)
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(cfg.KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("error creating partition consumer", err, partitionConsumer)

	}
	defer partitionConsumer.Close()
	fmt.Println("kafka consumer started")
	for {
		fmt.Println("started listening for new notifications")
		select {
		case message := <-partitionConsumer.Messages():
			msg, err := c.UnmarshelNotification(message.Value)
			if err != nil {
				fmt.Println("error unmarshelling message", err)
				continue
			}
			fmt.Println("received message", msg)
			err = c.notiRepository.StoreNotificationReq(*msg)

			if err != nil {
				fmt.Println("error storing notification in repository", err)
				continue
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("kafka cosumer error", err)

		}
	}
}

func (c *notificationUsecase) UnmarshelNotification(data []byte) (*models.NotificationReq, error) {
	var notification models.NotificationReq
	err := json.Unmarshal(data, &notification)
	if err != nil {
		return nil, err
	}
	notification.CreatedAt = time.Now()

	return &notification, nil
}

func (c *notificationUsecase) GetNotification(userid int, mod models.Pagination) ([]models.NotificationResponse, error) {
	data, err := c.notiRepository.GetNotification(userid, mod)
	if err != nil {
		return []models.NotificationResponse{}, err
	}
	var response []models.NotificationResponse
	for _, v := range data {
		// userdata, err := c.authclient.UserData(v.SenderID)
		// if err != nil {
		// 	return nil, err
		// }
		response = append(response, models.NotificationResponse{
			ID:        v.ID,
			UserID:    v.SenderID,
			Username:  v.SenderName,
			PostID:    v.PostID,
			Message:   v.Message,
			CreatedAt: v.CreatedAt.String(),
		})
	}
	return response, nil
}

func (c *notificationUsecase) ReadNotification(id, user_id int) (bool, error) {
	c.Logger.Info("ReadNotification at notificationUsecase started")
	if id <= 0 {
		c.Logger.Error("Error at notificationUsecase : ", errors.New("invalid notification id"))
		return false, errors.New("invalid notification id")
	}

	c.Logger.Info("IsNotificationExistOnUser at notiRepository started")
	ok, err := c.notiRepository.IsNotificationExistOnUser(id, user_id)
	if err != nil {
		c.Logger.Error("Error at IsNotificationExistOnUser at notiRepository: ", err)
		return false, err
	}
	if !ok {
		c.Logger.Error("Error at notificationUsecase : ", errors.New("notification not found"))
		return false, errors.New("notification not found")
	}
	c.Logger.Info("ReadNotification at notiRepository finished")
	c.Logger.Info("ReadNotification at notiRepository started")

	Ok, err := c.notiRepository.ReadNotification(id)

	if err != nil {
		c.Logger.Error("Error at ReadNotification at notiRepository: ", err)
		c.Logger.Error("Error at ReadNotification at notiRepository: ", err)
		return false, err
	}

	c.Logger.Info("ReadNotification at notiRepository finished")
	c.Logger.Info("ReadNotification at notificationUsecase finished")

	return Ok, nil

}

func (c *notificationUsecase) MarkAllAsRead(userId int) (bool, error) {
	c.Logger.Info("MarkAllAsRead at notificationUsecase started")

	c.Logger.Info("UnreadedNotificationExist at notiRepository started")
	ok, err := c.notiRepository.UnreadedNotificationExist(userId)
	if err != nil {
		c.Logger.Error("Error at UnreadedNotificationExist at notiRepository: ", err)
		return false, err
	}
	if !ok {
		c.Logger.Error("Error at notificationUsecase : ", errors.New("notification not found"))
		return false, errors.New("notifications not found")
	}
	c.Logger.Info("UnreadedNotificationExist at notiRepository finished")
	c.Logger.Info("MarkAllAsRead at notiRepository started")

	Ok, err := c.notiRepository.MarkAllAsRead(userId)

	if err != nil {
		c.Logger.Error("Error at MarkAllAsRead at notiRepository: ", err)
		return false, err
	}

	c.Logger.Info("MarkAllAsRead at notiRepository finished")
	c.Logger.Info("MarkAllAsRead at notificationUsecase finished")

	return Ok, nil

}

func (c *notificationUsecase) GetAllNotifications(userId int) ([]models.AllNotificationResponse, error) {
	c.Logger.Info("GetAllNotifications at notificationUsecase started")

	c.Logger.Info("GetAllNotifications at notiRepository started")

	data, err := c.notiRepository.GetAllNotifications(userId)

	if err != nil {
		c.Logger.Error("Error at MarkAllAsRead at notiRepository: ", err)
		return nil, err
	}

	c.Logger.Info("GetAllNotifications at notiRepository finished")
	c.Logger.Info("GetAllNotifications at notificationUsecase finished")

	return data, nil

}
