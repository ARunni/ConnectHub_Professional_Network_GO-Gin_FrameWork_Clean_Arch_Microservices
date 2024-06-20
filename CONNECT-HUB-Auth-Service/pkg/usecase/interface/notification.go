package interfaces

import "github.com/ARunni/ConnetHub_auth/pkg/utils/models"

type NotificationUsecase interface {
	UserData(userId int) (models.UserDatas, error)
}
