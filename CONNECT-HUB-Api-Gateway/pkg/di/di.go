package di

import (
	server "connectHub_gateway/pkg/api"
	"connectHub_gateway/pkg/api/handler"
	"connectHub_gateway/pkg/client"
	"connectHub_gateway/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	jobseekerClient := client.NewJobSeekerClient(cfg)
	jobseekerHandler := handler.NewJobSeekerHandler(jobseekerClient)

	recruiterClient := client.NewRecruiterClient(cfg)
	recruiterHandler := handler.NewRecruiterHandler(recruiterClient)

	jobClient := client.NewJobClient(cfg)
	jobHandler := handler.NewJobHandler(jobClient)

	serverHTTP := server.NewServerHTTP(adminHandler, jobseekerHandler, recruiterHandler, jobHandler)
	return serverHTTP, nil
}
