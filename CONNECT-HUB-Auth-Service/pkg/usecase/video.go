package usecase

import (
	"errors"
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	"github.com/ARunni/ConnetHub_auth/pkg/helper"
	repo "github.com/ARunni/ConnetHub_auth/pkg/repository/interface"
	usecase "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"

	"github.com/sirupsen/logrus"
)

type videoCallUseCase struct {
	videoCallRepository repo.VideoCallRepository
	Logger              *logrus.Logger
	LogFile             *os.File
}

func NewVideoCallUseCase(repo repo.VideoCallRepository) usecase.VideoCallUsecase {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &videoCallUseCase{
		videoCallRepository: repo,
		Logger:              logger,
		LogFile:             logFile,
	}
}

func (ur *videoCallUseCase) VideoCallKey(userID, oppositeUser int) (string, error) {

	ur.Logger.Info("usecase VideoCallKey function started")
	userExist, err := ur.videoCallRepository.IsRecruiterExist(userID)
	if err != nil {
		ur.Logger.Error("error from repo", err)
		return "", err
	}
	if !userExist {
		ur.Logger.Error("error from repo", errors.New("recruiter doesn't exist"))
		return "", errors.New("recruiter doesn't exist")
	}

	JobseekerExist, err := ur.videoCallRepository.IsJobseekerExist(oppositeUser)
	if err != nil {
		ur.Logger.Error("error from repo", err)
		return "", err
	}
	if !JobseekerExist {
		ur.Logger.Error("error from repo", errors.New("jobseeker doesn't exist"))
		return "", errors.New("user not found")
	}

	key, err := helper.GenerateVideoCallKey(userID, oppositeUser)
	if err != nil {
		ur.Logger.Error("error from helper", err)
		return "", err
	}
	ur.Logger.Info("usecase VideoCallKey function scuccess")
	return key, nil
}
