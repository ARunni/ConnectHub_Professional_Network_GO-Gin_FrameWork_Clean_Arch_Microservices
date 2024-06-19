package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	logging "github.com/ARunni/ConnetHub_Notification/Logging"
	"github.com/ARunni/ConnetHub_Notification/pkg/config"
	pb "github.com/ARunni/ConnetHub_Notification/pkg/pb/notification"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	listner net.Listener
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewGRPCServer(cfg config.Config, server pb.NotificationServiceServer) (*Server, error) {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Notification.log")
	logger.Info("NewGRPCServer started")
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		logger.Error("Failed to listen: ", err)
		return nil, err
	}
	logger.Info("NewGRPCServer success")

	newServer := grpc.NewServer(grpc.UnaryInterceptor(grpcInterceptor))

	pb.RegisterNotificationServiceServer(newServer, server)
	return &Server{
		server:  newServer,
		listner: lis,
		Logger:  logger,
		LogFile: logFile,
	}, nil
}

func (c *Server) Start() error {
	c.Logger.Info("grpc server listening on port :7006")
	fmt.Println("grpc server listening on 7006")
	return c.server.Serve(c.listner)
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
