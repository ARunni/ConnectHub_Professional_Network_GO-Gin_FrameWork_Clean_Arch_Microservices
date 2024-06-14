package repository

import (
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/repository/interface"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type videoCallRepository struct {
	DB      *gorm.DB
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewVideoCallRepository(DB *gorm.DB) interfaces.VideoCallRepository {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &videoCallRepository{
		DB:      DB,
		Logger:  logger,
		LogFile: logFile,
	}
}

func (jr *videoCallRepository) IsJobseekerExist(userId int) (bool, error) {
	var count int
	querry := ` select count(*) from job_seekers where id = ?`
	result := jr.DB.Raw(querry, userId).Scan(&count)
	if result.Error != nil {
		jr.Logger.Error("Error during database operations",result.Error)
		return false, result.Error
	}
	return count > 0, nil
}
func (jr *videoCallRepository) IsRecruiterExist(userId int) (bool, error) {
	var count int
	querry := ` select count(*) from recruiters where id = ?`
	result := jr.DB.Raw(querry, userId).Scan(&count)
	if result.Error != nil {
		jr.Logger.Error("Error during database operations",result.Error)
		return false, result.Error
	}
	return count > 0, nil
}
