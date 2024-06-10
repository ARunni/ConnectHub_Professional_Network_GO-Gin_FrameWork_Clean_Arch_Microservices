package usecase

import (
	logging "ConnetHub_post/Logging"
	repo "ConnetHub_post/pkg/repository/interfaces"
	"ConnetHub_post/pkg/usecase/interfaces"
	"os"

	"github.com/sirupsen/logrus"
)

type recruiterJobUseCase struct {
	postRepository repo.RecruiterPostRepository
	Logger         *logrus.Logger
	LogFile        *os.File
}

func NewRecruiterPostUseCase(repo repo.RecruiterPostRepository) interfaces.RecruiterPostUsecase {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Post.log")
	return &recruiterJobUseCase{
		postRepository: repo,
		Logger:         logger,
		LogFile:        logFile,
	}
}
