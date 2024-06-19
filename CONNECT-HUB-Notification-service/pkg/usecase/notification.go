package usecase

import (
	"encoding/json"
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

	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{cfg.KafkaBrokers}, configs)
	if err != nil {
		fmt.Println("error creating kafka consumer", err)

	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(cfg.KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("error creating partition consumer", err)

	}
	defer partitionConsumer.Close()
	fmt.Println("kafka consumer started")
	for {
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
	fmt.Println("dddddddd", data)
	for _, v := range data {
		userdata, err := c.authclient.UserData(v.SenderID)
		if err != nil {
			fmt.Println("heloooooooooooo")
			return nil, err
		}
		response = append(response, models.NotificationResponse{
			UserID:    int(userdata.UserId),
			Username:  userdata.Username,
			Profile:   userdata.Profile,
			Message:   v.Message,
			CreatedAt: v.CreatedAt.String(),
		})
	}
	return response, nil
}
