package client

import (
	logging "ConnetHub_job/Logging"
	"ConnetHub_job/pkg/client/auth/interfaces"
	"ConnetHub_job/pkg/config"
	pb "ConnetHub_job/pkg/pb/auth"
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type JobAuthClient struct {
	Client  pb.JobAuthClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewJobAuthClient(cfg config.Config) interfaces.JobAuthClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Job.log")
	grpcConnection, err := grpc.Dial(cfg.Connect_Hub_Auth, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth jobseeker", err)
	}

	grpcClient := pb.NewJobAuthClient(grpcConnection)
	return &JobAuthClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}

}

func (jc *JobAuthClient) GetDetailsById(userId int) (string, string, error) {
	fmt.Println("herre client job Auth", userId)
	data, err := jc.Client.GetDetailsById(context.Background(), &pb.GetDetailsByIdRequest{
		Userid: int64(userId),
	})
	if err != nil {
		return "", "", err
	}
	return data.Email, data.Username, nil
}
