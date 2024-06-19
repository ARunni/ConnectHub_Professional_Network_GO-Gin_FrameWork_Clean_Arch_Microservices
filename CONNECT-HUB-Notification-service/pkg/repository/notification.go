package repository

import (
	"os"

	logging "github.com/ARunni/ConnetHub_Notification/Logging"
	interfaces "github.com/ARunni/ConnetHub_Notification/pkg/repository/interface"
	"github.com/ARunni/ConnetHub_Notification/pkg/utils/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type notificationRepository struct {
	DB      *gorm.DB
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewNotificationRepository(DB *gorm.DB) interfaces.NotificationRepository {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Notification.log")
	return &notificationRepository{
		DB:      DB,
		Logger:  logger,
		LogFile: logFile,
	}
}

func (c *notificationRepository) StoreNotificationReq(noti models.NotificationReq) error {
	c.Logger.Info("StoreNotificationReq at notificationRepository started")
	err := c.DB.Exec("INSERT INTO notifications(user_id,sender_id,post_id,message,created_at) VALUES(?,?,?,?,?)", noti.UserID, noti.SenderID, noti.PostID, noti.Message, noti.CreatedAt).Error
	if err != nil {
		c.Logger.Error("Error while storing notification request: ", err)
		return err
	}
	c.Logger.Info("StoreNotificationReq at notificationRepository completed successfully")
	return nil
}

func (c *notificationRepository) GetNotification(id int, pag models.Pagination) ([]models.Notification, error) {
	c.Logger.Info("GetNotification at notificationRepository started")

	var data []models.Notification
	if pag.Offset <= 0 {
		pag.Offset = 1
	}
	offset := (pag.Offset - 1) * pag.Limit
	err := c.DB.Raw("SELECT sender_id,message, created_at FROM notifications WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", id, pag.Limit, offset).Scan(&data).Error
	if err != nil {
		c.Logger.Error("Error while getting notifications: ", err)
		return nil, err
	}
	c.Logger.Info("GetNotification at notificationRepository completed successfully")
	return data, nil
}
