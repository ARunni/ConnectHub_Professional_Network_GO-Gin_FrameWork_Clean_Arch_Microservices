package usecase

import (
	logging "github.com/ARunni/ConnetHub_post/Logging"
	repo "github.com/ARunni/ConnetHub_post/pkg/repository/interfaces"
	"github.com/ARunni/ConnetHub_post/pkg/usecase/interfaces"
	"os"

	"github.com/sirupsen/logrus"
)

type adminJobUseCase struct {
	postRepository repo.AdminPostRepository
	Logger         *logrus.Logger
	LogFile        *os.File
}

func NewAdminPostUseCase(repo repo.AdminPostRepository) interfaces.AdminPostUsecase {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Post.log")
	return &adminJobUseCase{
		postRepository: repo,
		Logger:         logger,
		LogFile:        logFile,
	}
}
