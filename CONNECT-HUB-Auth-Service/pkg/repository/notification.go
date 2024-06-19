package repository

import (
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/repository/interface"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type notificationRepository struct {
	DB      *gorm.DB
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewNotificationRepository(DB *gorm.DB) interfaces.NotificationRepository {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &notificationRepository{
		DB:      DB,
		Logger:  logger,
		LogFile: logFile,
	}
}
