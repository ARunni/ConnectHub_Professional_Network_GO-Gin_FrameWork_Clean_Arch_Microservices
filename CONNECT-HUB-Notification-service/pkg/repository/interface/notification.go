package interfaces

import "github.com/ARunni/ConnetHub_Notification/pkg/utils/models"

type NotificationRepository interface {
	StoreNotificationReq(models.NotificationReq) error
	GetNotification(userid int, req models.Pagination) ([]models.Notification, error)
	ReadNotification(id int) (bool,error)
	IsNotificationExistOnUser(id,userId int) (bool,error)
	MarkAllAsRead(userId int) (bool,error)
	UnreadedNotificationExist(userId int) (bool,error)
}
