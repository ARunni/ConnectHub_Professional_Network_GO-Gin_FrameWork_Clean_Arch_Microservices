package repository

import (
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/repository/interface"
	"github.com/ARunni/ConnetHub_auth/pkg/utils/models"
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

func (nr *notificationRepository) UserData(userId int) (models.UserDatas, error) {
	nr.Logger.Info("UserData at notificationRepository started")
	var data models.UserDatas
	querry := `
select first_name as username,id as user_id from users where id = ?
`
	result := nr.DB.Raw(querry, userId).Scan(&data)
	if result.Error != nil {
		nr.Logger.Error("error on database", result.Error)
		return models.UserDatas{}, result.Error
	}
	nr.Logger.Info("UserData at notificationRepository finished")
	return data, nil

}
