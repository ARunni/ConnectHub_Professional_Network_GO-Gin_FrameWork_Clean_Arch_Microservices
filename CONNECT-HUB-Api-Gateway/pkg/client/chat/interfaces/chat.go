package interfaces

import "connectHub_gateway/pkg/utils/models"

type ChatClient interface {
	GetChat(userID string, req models.ChatRequest) ([]models.TempMessage, error)
}
