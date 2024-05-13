package di

import (
	server "connectHub_gateway/pkg/api"
	"connectHub_gateway/pkg/api/handler"
	authClient "connectHub_gateway/pkg/client/auth"
	jobClient "connectHub_gateway/pkg/client/job"
	"connectHub_gateway/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminAuthClient := authClient.NewAdminAuthClient(cfg)
	adminAuthHandler := handler.NewAdminAuthHandler(adminAuthClient)

	jobseekerAuthClient := authClient.NewJobSeekerAuthClient(cfg)
	jobseekerAuthHandler := handler.NewJobSeekerAuthHandler(jobseekerAuthClient)

	recruiterAuthClient := authClient.NewRecruiterAuthClient(cfg)
	recruiterAuthHandler := handler.NewRecruiterAuthHandler(recruiterAuthClient)

	recruiterJobClient := jobClient.NewRecruiterJobClient(cfg)
	recruiterJobHandler := handler.NewRecruiterJobHandler(recruiterJobClient)

	serverHTTP := server.NewServerHTTP(adminAuthHandler, jobseekerAuthHandler, recruiterAuthHandler, recruiterJobHandler)
	return serverHTTP, nil
}
