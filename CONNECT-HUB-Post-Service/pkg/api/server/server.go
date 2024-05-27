package server

import (
	"ConnetHub_post/pkg/config"
	postJ"ConnetHub_post/pkg/pb/post/jobseeker"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, JPostServer postJ.JobseekerPostServiceServer   ) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer(grpc.UnaryInterceptor(grpcInterceptor))
	postJ.RegisterJobseekerPostServiceServer(newServer, JPostServer)
	// jobJ.RegisterJobseekerJobServer(newServer, JJobServer)

	// recruiterPb.RegisterRecruiterServer(newServer, recruiterServer)
	// jobseekerPb.RegisterJobseekerServer(newServer, jobseekerServer)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (s *Server) Start() error {
	fmt.Println("grpc server listening on port :7004")
	return s.server.Serve(s.listener)
}

func grpcInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf(color.GreenString("Received gRPC request: %s"), info.FullMethod)
	// Call the handler function to process the request
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf(color.RedString("gRPC request failed: %v"), err)
	} else {
		log.Printf(color.GreenString("Completed gRPC request: %s"), info.FullMethod)
	}
	return resp, err
}
