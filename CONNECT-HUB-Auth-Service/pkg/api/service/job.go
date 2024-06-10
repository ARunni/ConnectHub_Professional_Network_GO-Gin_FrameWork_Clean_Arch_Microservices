package service

import (
	logging "github.com/ARunni/ConnetHub_auth/Logging"
	pb "github.com/ARunni/ConnetHub_auth/pkg/pb/job/auth"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type JobAuthServer struct {
	recruiterUsecase interfaces.RecruiterUseCase
	pb.UnimplementedJobAuthServer
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewJobauthServer(usecase interfaces.RecruiterUseCase) pb.JobAuthServer {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &JobAuthServer{
		recruiterUsecase: usecase,
		Logger:           logger,
		LogFile:          logFile,
	}
}

func (js *JobAuthServer) GetDetailsById(ctx context.Context, Req *pb.GetDetailsByIdRequest) (*pb.GetDetailsByIdResponse, error) {
	email, name, err := js.recruiterUsecase.GetDetailsById(int(Req.Userid))
	if err != nil {
		return nil, err
	}
	return &pb.GetDetailsByIdResponse{
		Username: name,
		Email:    email,
	}, nil
}
