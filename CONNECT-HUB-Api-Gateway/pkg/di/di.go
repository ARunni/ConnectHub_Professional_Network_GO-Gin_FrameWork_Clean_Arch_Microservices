package di

import (
	server "connectHub_gateway/pkg/api"
	authHandler "connectHub_gateway/pkg/api/handler/auth"
	chatHandler "connectHub_gateway/pkg/api/handler/chat"
	jobHandler "connectHub_gateway/pkg/api/handler/job"
	postHandler "connectHub_gateway/pkg/api/handler/post"
	"connectHub_gateway/pkg/helper"

	authClient "connectHub_gateway/pkg/client/auth"
	chatClient "connectHub_gateway/pkg/client/chat"
	jobClient "connectHub_gateway/pkg/client/job"
	postClient "connectHub_gateway/pkg/client/post"

	"connectHub_gateway/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminAuthClient := authClient.NewAdminAuthClient(cfg)
	adminAuthHandler := authHandler.NewAdminAuthHandler(adminAuthClient)

	jobseekerAuthClient := authClient.NewJobSeekerAuthClient(cfg)
	jobseekerAuthHandler := authHandler.NewJobSeekerAuthHandler(jobseekerAuthClient)

	recruiterAuthClient := authClient.NewRecruiterAuthClient(cfg)
	recruiterAuthHandler := authHandler.NewRecruiterAuthHandler(recruiterAuthClient)

	recruiterJobClient := jobClient.NewRecruiterJobClient(cfg)
	recruiterJobHandler := jobHandler.NewRecruiterJobHandler(recruiterJobClient)

	jobseekerJobClient := jobClient.NewJobseekerJobClient(cfg)
	JobseekerJobhandler := jobHandler.NewJobseekerJobHandler(jobseekerJobClient)

	jobseekerPostClient := postClient.NewJobseekerPostClient(cfg)
	jobseekerPostHandler := postHandler.NewJobseekerPostHandler(jobseekerPostClient)


	chatClient := chatClient.NewChatClient(cfg)
	chatHandler := chatHandler.NewChatHandler(chatClient,helper.NewHelper(&cfg))

	serverHTTP := server.NewServerHTTP(
		adminAuthHandler, jobseekerAuthHandler,
		recruiterAuthHandler, recruiterJobHandler,
		JobseekerJobhandler,jobseekerPostHandler,
		chatHandler)

	return serverHTTP, nil
}
