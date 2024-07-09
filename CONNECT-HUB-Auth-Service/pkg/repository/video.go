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
	jr.Logger.Info("repo IsJobseekerExist started ")
	var count int
	querry := ` select count(*) from users where id = ? and role = ?`
	result := jr.DB.Raw(querry, userId,"Jobseeker").Scan(&count)
	if result.Error != nil {
		jr.Logger.Error("repo IsJobseekerExist error ", result.Error)
		return false, result.Error
	}
	jr.Logger.Info("repo IsJobseekerExist success ")
	return count > 0, nil
}
// func (jr *videoCallRepository) IsJobseekerExist(userId int) (bool, error) {
// 	jr.Logger.Info("repo IsJobseekerExist started ")
// 	var count int
// 	querry := ` select count(*) from job_seekers where id = ?`
// 	result := jr.DB.Raw(querry, userId).Scan(&count)
// 	if result.Error != nil {
// 		jr.Logger.Error("repo IsJobseekerExist error ", result.Error)
// 		return false, result.Error
// 	}
// 	jr.Logger.Info("repo IsJobseekerExist success ")
// 	return count > 0, nil
// }
func (jr *videoCallRepository) IsRecruiterExist(userId int) (bool, error) {
	jr.Logger.Info("repo IsRecruiterExist started ")
	var count int
	querry := ` select count(*) from users where id = ? and role = ?`
	result := jr.DB.Raw(querry, userId,"Recruiter").Scan(&count)
	if result.Error != nil {
		jr.Logger.Error("repo IsRecruiterExist error ", result.Error)
		return false, result.Error
	}
	jr.Logger.Info("repo IsRecruiterExist success ")
	return count > 0, nil
}
// func (jr *videoCallRepository) IsRecruiterExist(userId int) (bool, error) {
// 	jr.Logger.Info("repo IsRecruiterExist started ")
// 	var count int
// 	querry := ` select count(*) from recruiters where id = ?`
// 	result := jr.DB.Raw(querry, userId).Scan(&count)
// 	if result.Error != nil {
// 		jr.Logger.Error("repo IsRecruiterExist error ", result.Error)
// 		return false, result.Error
// 	}
// 	jr.Logger.Info("repo IsRecruiterExist success ")
// 	return count > 0, nil
// }
