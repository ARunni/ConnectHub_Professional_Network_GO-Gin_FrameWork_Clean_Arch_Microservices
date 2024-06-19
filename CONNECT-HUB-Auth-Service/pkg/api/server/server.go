package server

import (
	"context"
	"log"
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	"github.com/ARunni/ConnetHub_auth/pkg/config"
	adminPb "github.com/ARunni/ConnetHub_auth/pkg/pb/auth/admin"

	"fmt"
	"net"

	authPb "github.com/ARunni/ConnetHub_auth/pkg/pb/auth/auth"
	jobseekerPb "github.com/ARunni/ConnetHub_auth/pkg/pb/auth/jobseeker"
	recruiterPb "github.com/ARunni/ConnetHub_auth/pkg/pb/auth/recruiter"
	jobAuthPb "github.com/ARunni/ConnetHub_auth/pkg/pb/job/auth"
	notificationauthPb "github.com/ARunni/ConnetHub_auth/pkg/pb/notification/auth"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
	Logger   *logrus.Logger
	LogFile  *os.File
}

func NewGRPCServer(cfg config.Config, adminServer adminPb.AdminServer,
	recruiterServer recruiterPb.RecruiterServer,
	jobseekerServer jobseekerPb.JobseekerServer,
	jobAuthServer jobAuthPb.JobAuthServer, authServer authPb.AuthServiceServer,
	notificationAuthserver notificationauthPb.NotificationAuthServiceServer) (*Server, error) {

	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	logger.Info("NewGRPCServer started")
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		logger.Error("Failed to listen: ", err)
		return nil, err
	}

	newServer := grpc.NewServer(grpc.UnaryInterceptor(grpcInterceptor))
	adminPb.RegisterAdminServer(newServer, adminServer)
	recruiterPb.RegisterRecruiterServer(newServer, recruiterServer)
	jobseekerPb.RegisterJobseekerServer(newServer, jobseekerServer)
	jobAuthPb.RegisterJobAuthServer(newServer, jobAuthServer)
	authPb.RegisterAuthServiceServer(newServer, authServer)
	notificationauthPb.RegisterNotificationAuthServiceServer(newServer, notificationAuthserver)

	logger.Info("NewGRPCServer success")
	return &Server{
		server:   newServer,
		listener: lis,
		Logger:   logger,
		LogFile:  logFile,
	}, nil
}

func (s *Server) Start() error {
	s.Logger.Info("grpc server listening on port :7001")
	fmt.Println("grpc server listening on port :7001")
	return s.server.Serve(s.listener)
}

func grpcInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf(color.GreenString("Received gRPC request: %s"), info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf(color.RedString("gRPC request failed: %v"), err)
	} else {
		log.Printf(color.GreenString("Completed gRPC request: %s"), info.FullMethod)
	}
	return resp, err
}
