package server

import (
	"ConnetHub_auth/pkg/config"
	adminPb "ConnetHub_auth/pkg/pb/auth/admin"

	// recruiterPb "ConnetHub_auth/pkg/pb/auth/jobseeker"
	// jobseekerPb "ConnetHub_auth/pkg/pb/auth/recruiter"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, adminServer adminPb.AdminServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	adminPb.RegisterAdminServer(newServer, adminServer)
	// pb.RegisterEmployerServer(newServer, recruiterServer)
	// pb.RegisterJobSeekerServer(newServer, jobseekerServer)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50052")
	return c.server.Serve(c.listener)
}
