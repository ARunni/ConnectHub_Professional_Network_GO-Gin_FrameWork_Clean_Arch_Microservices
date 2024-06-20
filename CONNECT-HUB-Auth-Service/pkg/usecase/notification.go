package usecase

import (
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	repo "github.com/ARunni/ConnetHub_auth/pkg/repository/interface"
	usecase "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"
	"github.com/ARunni/ConnetHub_auth/pkg/utils/models"
	"github.com/sirupsen/logrus"
)

type notificationUseCase struct {
	notificationRepository repo.NotificationRepository
	Logger                 *logrus.Logger
	LogFile                *os.File
}

func NewNotificationUseCase(repo repo.NotificationRepository) usecase.NotificationUsecase {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &notificationUseCase{
		notificationRepository: repo,
		Logger:                 logger,
		LogFile:                logFile,
	}
}

func (nu *notificationUseCase) UserData(userId int) (models.UserDatas, error) {
	nu.Logger.Info("UserData at notificationUseCase started")
	nu.Logger.Info("UserData at notificationRepository started")
	data, err := nu.notificationRepository.UserData(userId)
	if err != nil {
		nu.Logger.Error("error at notificationRepository", err)
		return models.UserDatas{}, err
	}
	nu.Logger.Info("UserData at notificationRepository finished")
	nu.Logger.Info("UserData at notificationUseCase finished")
	return data, nil
}
