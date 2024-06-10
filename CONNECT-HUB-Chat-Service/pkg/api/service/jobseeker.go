package service

import (
	pb "github.com/ARunni/ConnetHub_chat/pkg/pb/chat"
	interfaces "github.com/ARunni/ConnetHub_chat/pkg/usecase/interface"
	"github.com/ARunni/ConnetHub_chat/pkg/utils/models"
	"context"
	"time"
)

type ChatServer struct {
	chatUseCase interfaces.ChatUseCase
	pb.UnimplementedChatServiceServer
}

func NewChatServer(UseCaseChat interfaces.ChatUseCase) pb.ChatServiceServer {
	return &ChatServer{
		chatUseCase: UseCaseChat,
	}
}

func (c *ChatServer) GetFriendChat(ctx context.Context, req *pb.GetFriendChatRequest) (*pb.GetFriendChatResponse, error) {
	ind, _ := time.LoadLocation("Asia/Kolkata")
	result, err := c.chatUseCase.GetFriendChat(req.UserID, req.FriendID, models.Pagination{Limit: req.Limit, OffSet: req.OffSet})
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.Message
	for _, val := range result {
		finalResult = append(finalResult, &pb.Message{
			MessageID:   val.ID,
			SenderId:    val.SenderID,
			RecipientId: val.RecipientID,
			Content:     val.Content,
			Timestamp:   val.Timestamp.In(ind).String(),
		})
	}
	return &pb.GetFriendChatResponse{FriendChat: finalResult}, nil
}
