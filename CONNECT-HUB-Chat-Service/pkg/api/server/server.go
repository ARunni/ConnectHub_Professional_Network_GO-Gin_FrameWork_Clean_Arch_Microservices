package server

import (
	"github.com/ARunni/ConnetHub_chat/pkg/config"
	"fmt"
	"net"

	pb "github.com/ARunni/ConnetHub_chat/pkg/pb/chat"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.ChatServiceServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	newServer := grpc.NewServer()
	pb.RegisterChatServiceServer(newServer, server)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :7003")
	return c.server.Serve(c.listener)
}
