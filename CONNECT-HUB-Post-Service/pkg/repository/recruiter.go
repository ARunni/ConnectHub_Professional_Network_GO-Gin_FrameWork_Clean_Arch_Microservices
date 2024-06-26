package repository

import (
	logging "github.com/ARunni/ConnetHub_post/Logging"
	"github.com/ARunni/ConnetHub_post/pkg/repository/interfaces"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type recruiterPostRepository struct {
	DB      *gorm.DB
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewRecruiterPostRepository(DB *gorm.DB) interfaces.RecruiterPostRepository {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Post.log")
	return &recruiterPostRepository{
		DB:      DB,
		Logger:  logger,
		LogFile: logFile,
	}
}
