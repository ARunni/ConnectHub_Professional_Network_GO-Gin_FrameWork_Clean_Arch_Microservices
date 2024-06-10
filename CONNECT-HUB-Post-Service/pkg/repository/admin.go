package repository

import (
	logging "github.com/ARunni/ConnetHub_post/Logging"
	"github.com/ARunni/ConnetHub_post/pkg/repository/interfaces"
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
