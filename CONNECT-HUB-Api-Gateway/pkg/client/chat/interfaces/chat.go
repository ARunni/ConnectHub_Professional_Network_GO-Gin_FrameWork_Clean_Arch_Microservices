package interfaces

import "github.com/ARunni/connectHub_gateway/pkg/utils/models"

type ChatClient interface {
	GetChat(userID string, req models.ChatRequest) ([]models.TempMessage, error)
}
