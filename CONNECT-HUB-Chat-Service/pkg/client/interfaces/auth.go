package interfaces

import "ConnetHub_chat/pkg/utils/models"

type NewauthClient interface {
	CheckUserAvalilabilityWithUserID(userID int) bool
	UserData(userID int) (models.UserData, error)
}
