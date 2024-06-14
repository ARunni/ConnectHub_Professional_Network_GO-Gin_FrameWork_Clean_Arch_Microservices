package service

import (
	"context"
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	pb "github.com/ARunni/ConnetHub_auth/pkg/pb/auth/auth"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"

	"github.com/sirupsen/logrus"
)

type authServer struct {
	authUsecase interfaces.VideoCallUsecase
	pb.UnimplementedAuthServiceServer
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewauthServer(useCase interfaces.VideoCallUsecase) pb.AuthServiceServer {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &authServer{
		authUsecase: useCase,
		Logger:      logger,
		LogFile:     logFile,
	}
}

func (au *authServer) VideoCallKey(ctx context.Context, req *pb.VideoCallRequest) (*pb.VideoCallResponse, error) {
	au.Logger.Info("authserver VideoCallKey started")
	key, err := au.authUsecase.VideoCallKey(int(req.UserID), int(req.OppositeUser))
	if err != nil {
		au.Logger.Error("error from usecase VideoCallKey", err)
		return &pb.VideoCallResponse{}, err
	}
	au.Logger.Info("success from usecase VideoCallKey")
	return &pb.VideoCallResponse{
		Key: key,
	}, nil
}
