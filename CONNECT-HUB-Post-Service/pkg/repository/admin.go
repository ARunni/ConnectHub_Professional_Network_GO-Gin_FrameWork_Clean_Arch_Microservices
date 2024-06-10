package repository

import (
	logging "ConnetHub_post/Logging"
	"ConnetHub_post/pkg/repository/interfaces"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type adminPostRepository struct {
	DB      *gorm.DB
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewAdminPostRepository(DB *gorm.DB) interfaces.AdminPostRepository {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Post.log")
	return &adminPostRepository{
		DB:      DB,
		Logger:  logger,
		LogFile: logFile,
	}
}
