package client

import (
	"os"

	logging "github.com/ARunni/ConnetHub_chat/Logging"
	"github.com/ARunni/ConnetHub_chat/pkg/config"
	"github.com/ARunni/ConnetHub_chat/pkg/utils/models"

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
		logger.Error("Could not connect grpc")
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
	c.Logger.Info("CheckUserAvalilabilityWithUserID at clientAuth started")
	c.Logger.Info("CheckUserAvalilabilityWithUserID at Client started")
	ok, _ := c.Client.CheckUserAvalilabilityWithUserID(context.Background(), &pb.CheckUserAvalilabilityWithUserIDRequest{
		Id: int64(userID),
	})
	c.Logger.Info("CheckUserAvalilabilityWithUserID at Client finished")
	c.Logger.Info("CheckUserAvalilabilityWithUserID at clientAuth finished")
	return ok.Valid
}

func (c *clientAuth) UserData(userID int) (models.UserData, error) {
	c.Logger.Info("UserData at clientAuth started")
	c.Logger.Info("UserData at Client started")
	data, err := c.Client.UserData(context.Background(), &pb.UserDataRequest{
		Id: int64(userID),
	})
	if err != nil {
		c.Logger.Error("error from client", err)
		return models.UserData{}, err
	}
	c.Logger.Info("UserData at Client finished")
	c.Logger.Info("UserData at clientAuth finished")
	return models.UserData{
		UserId:   uint(data.Id),
		Username: data.Username,
		Profile:  data.ProfilePhoto,
	}, nil
}
