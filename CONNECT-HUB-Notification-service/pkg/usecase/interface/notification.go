package interfaces

import "github.com/ARunni/ConnetHub_Notification/pkg/utils/models"

type NotificationUseCase interface {
	ConsumeNotification()
	GetNotification(userid int, mod models.Pagination) ([]models.NotificationResponse, error)
}
