package client

import (
	"context"
	"fmt"
	"os"

	logging "github.com/ARunni/ConnetHub_Notification/Logging"
	"github.com/ARunni/ConnetHub_Notification/pkg/config"
	pb "github.com/ARunni/ConnetHub_Notification/pkg/pb/auth"
	"github.com/ARunni/ConnetHub_Notification/pkg/utils/models"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type authClient struct {
	Client  pb.AuthServiceClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewAuthClient(cfg *config.Config) *authClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Notification.log")
	grpcConnection, err := grpc.Dial(cfg.Connect_Hub_Auth, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect", err)
	}
	grpcClient := pb.NewAuthServiceClient(grpcConnection)

	return &authClient{
		Client: grpcClient,
		Logger:              logger,
		LogFile:             logFile,
	}
}

func (ad *authClient) UserData(userid int) (models.UserData, error) {
	ad.Logger.Info("UserData at authClient started")
	
	data, err := ad.Client.UserData(context.Background(), &pb.UserDataRequest{
		Userid: int64(userid),
	})

	if err != nil {
		ad.Logger.Error("error from authClient", err)
		return models.UserData{}, err
	}
	ad.Logger.Info("UserData at authClient success")
	return models.UserData{
		UserId:   int(data.Userid),
		Username: data.Username,
		Profile:  data.Profile,
	}, nil
}
