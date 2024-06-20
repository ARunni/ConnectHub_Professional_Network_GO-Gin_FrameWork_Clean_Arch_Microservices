package interfaces

import "github.com/ARunni/ConnetHub_auth/pkg/utils/models"

type NotificationRepository interface {
	UserData(userId int) (models.UserDatas, error)
}
