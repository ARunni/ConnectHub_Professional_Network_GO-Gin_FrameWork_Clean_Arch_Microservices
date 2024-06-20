package client

import (
	logging "github.com/ARunni/ConnetHub_job/Logging"
	"github.com/ARunni/ConnetHub_job/pkg/client/auth/interfaces"
	"github.com/ARunni/ConnetHub_job/pkg/config"
	pb "github.com/ARunni/ConnetHub_job/pkg/pb/auth"
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
	jc.Logger.Info("GetDetailsById at JobAuthClient started")
	jc.Logger.Info("GetDetailsById at Client started")
	data, err := jc.Client.GetDetailsById(context.Background(), &pb.GetDetailsByIdRequest{
		Userid: int64(userId),
	})
	if err != nil {
		jc.Logger.Error("error from grpc call GetDetailsById",err)
		return "", "", err
	}
	jc.Logger.Info("GetDetailsById at JobAuthClient success")
	jc.Logger.Info("GetDetailsById at Client success")
	return data.Email, data.Username, nil
}

func (jc *JobAuthClient) GetDetailsByIdRecuiter(userId int) (string, string, error) {
	jc.Logger.Info("GetDetailsByIdRecuiter at JobAuthClient started")
	jc.Logger.Info("GetDetailsByIdRecuiter at Client started")
	data, err := jc.Client.GetDetailsByIdRecuiter(context.Background(), &pb.GetDetailsByIdRequest{
		Userid: int64(userId),
	})
	if err != nil {
		jc.Logger.Error("error from grpc call GetDetailsByIdRecuiter",err)
		return "", "", err
	}
	jc.Logger.Info("GetDetailsByIdRecuiter at JobAuthClient success")
	jc.Logger.Info("GetDetailsByIdRecuiter at Client success")
	return data.Email, data.Username, nil
}
