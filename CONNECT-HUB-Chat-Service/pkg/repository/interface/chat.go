
package interfaces

import "github.com/ARunni/ConnetHub_chat/pkg/utils/models"

type ChatRepository interface {
	StoreFriendsChat(models.MessageReq) error
	GetFriendChat(string, string, models.Pagination) ([]models.Message, error)
	UpdateReadAsMessage(string, string) error
}
