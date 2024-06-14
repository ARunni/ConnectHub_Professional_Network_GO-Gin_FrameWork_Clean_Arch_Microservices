package client

import (
	logging "github.com/ARunni/connectHub_gateway/Logging"
	"github.com/ARunni/connectHub_gateway/pkg/config"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"
	"os"

	pb "github.com/ARunni/connectHub_gateway/pkg/pb/chat"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ChatClient struct {
	Client  pb.ChatServiceClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewChatClient(cfg config.Config) *ChatClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	grpcConnection, err := grpc.Dial(cfg.ConnetHubChat, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewChatServiceClient(grpcConnection)

	return &ChatClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}
}

func (c *ChatClient) GetChat(userID string, req models.ChatRequest) ([]models.TempMessage, error) {
	c.Logger.Info("GetChat at client started")
	data, err := c.Client.GetFriendChat(context.Background(), &pb.GetFriendChatRequest{
		UserID:   userID,
		FriendID: req.FriendID,
		OffSet:   req.Offset,
		Limit:    req.Limit,
	})
	if err != nil {
		c.Logger.Error("GetChat at client failed: ", err)
		return []models.TempMessage{}, err
	}
	var response []models.TempMessage
	for _, v := range data.FriendChat {
		chatResponse := models.TempMessage{
			SenderID:    v.SenderId,
			RecipientID: v.RecipientId,
			Content:     v.Content,
			Timestamp:   v.Timestamp,
		}
		response = append(response, chatResponse)

	}
	c.Logger.Info("GetChat at client finished")
	return response, nil
}
