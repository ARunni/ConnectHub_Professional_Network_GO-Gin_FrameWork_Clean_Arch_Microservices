package interfaces

import "github.com/ARunni/ConnetHub_Notification/pkg/utils/models"

type NotificationUseCase interface {
	ConsumeNotification()
	GetNotification(userid int, mod models.Pagination) ([]models.NotificationResponse, error)
	ReadNotification(id, user_id int) (bool, error)
	MarkAllAsRead(userId int) (bool,error)
}
