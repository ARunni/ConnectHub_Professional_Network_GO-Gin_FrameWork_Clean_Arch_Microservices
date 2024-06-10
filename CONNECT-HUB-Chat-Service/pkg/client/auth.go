package client

import (
	logging "github.com/ARunni/ConnetHub_chat/Logging"
	"github.com/ARunni/ConnetHub_chat/pkg/config"
	"github.com/ARunni/ConnetHub_chat/pkg/utils/models"
	"os"

	"context"
	"fmt"

	pb "github.com/ARunni/ConnetHub_chat/pkg/pb/auth"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type clientAuth struct {
	Client  pb.AuthServiceClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewAuthClient(c *config.Config) *clientAuth {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Chat.log")
	cc, err := grpc.Dial(c.Connect_Hub_Auth, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	pbClient := pb.NewAuthServiceClient(cc)

	return &clientAuth{
		Client:  pbClient,
		Logger:  logger,
		LogFile: logFile,
	}
}

func (c *clientAuth) CheckUserAvalilabilityWithUserID(userID int) bool {
	ok, _ := c.Client.CheckUserAvalilabilityWithUserID(context.Background(), &pb.CheckUserAvalilabilityWithUserIDRequest{
		Id: int64(userID),
	})
	return ok.Valid
}

func (c *clientAuth) UserData(userID int) (models.UserData, error) {
	data, err := c.Client.UserData(context.Background(), &pb.UserDataRequest{
		Id: int64(userID),
	})
	if err != nil {
		return models.UserData{}, err
	}
	return models.UserData{
		UserId:   uint(data.Id),
		Username: data.Username,
		Profile:  data.ProfilePhoto,
	}, nil
}
