package client

import (
	"context"
	"fmt"
	"os"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	"github.com/ARunni/connectHub_gateway/pkg/config"
	pb "github.com/ARunni/connectHub_gateway/pkg/pb/auth/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/auth/interface"
)

type authClient struct {
	Client  pb.AuthServiceClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewAuthClient(cfg config.Config) interfaces.AuthClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	grpcConnection, err := grpc.Dial(cfg.ConnetHubAuth, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth jobseeker", err)
	}

	grpcClient := pb.NewAuthServiceClient(grpcConnection)
	return &authClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}

}

func (au *authClient) VideoCallKey(userID, oppositeUser int) (string, error) {
	key, err := au.Client.VideoCallKey(context.Background(), &pb.VideoCallRequest{
		UserID:       int64(userID),
		OppositeUser: int64(oppositeUser),
	})
	if err != nil {
		return "", err
	}
	return key.Key, nil
}
