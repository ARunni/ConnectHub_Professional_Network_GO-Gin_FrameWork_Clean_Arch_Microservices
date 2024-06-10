package interfaces

import "github.com/ARunni/ConnetHub_chat/pkg/utils/models"

type ChatUseCase interface {
	MessageConsumer()
	GetFriendChat(string, string, models.Pagination) ([]models.Message, error)
}
