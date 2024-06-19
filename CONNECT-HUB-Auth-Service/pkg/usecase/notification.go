package usecase

import (
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	repo "github.com/ARunni/ConnetHub_auth/pkg/repository/interface"
	usecase "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"
	"github.com/sirupsen/logrus"
)

type notificationUseCase struct {
	jobseekerRepository repo.NotificationRepository
	Logger              *logrus.Logger
	LogFile             *os.File
}

func NewNotificationUseCase(repo repo.NotificationRepository) usecase.NotificationUsecase {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &notificationUseCase{
		jobseekerRepository: repo,
		Logger:              logger,
		LogFile:             logFile,
	}
}
