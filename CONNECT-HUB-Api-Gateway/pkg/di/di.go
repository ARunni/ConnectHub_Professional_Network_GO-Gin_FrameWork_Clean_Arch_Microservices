package di

import (
	server "github.com/ARunni/connectHub_gateway/pkg/api"
	videoHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/Video_Call"
	authHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/auth"
	chatHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/chat"
	jobHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/job"
	notificationHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/notification"
	postHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/post"
	"github.com/ARunni/connectHub_gateway/pkg/helper"

	authClient "github.com/ARunni/connectHub_gateway/pkg/client/auth"
	chatClient "github.com/ARunni/connectHub_gateway/pkg/client/chat"
	jobClient "github.com/ARunni/connectHub_gateway/pkg/client/job"
	notificationClient "github.com/ARunni/connectHub_gateway/pkg/client/notification"
	postClient "github.com/ARunni/connectHub_gateway/pkg/client/post"

	"github.com/ARunni/connectHub_gateway/pkg/config"
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

	notificationClient := notificationClient.NewNotificationClient(cfg)
	notificationHandler := notificationHandler.NewNotificationHandler(notificationClient)

	VideoHandler := videoHandler.NewVideoCallHandler()
	authclient := authClient.NewAuthClient(cfg)
	authHandler := authHandler.NewAuthHandler(authclient)

	chatClient := chatClient.NewChatClient(cfg)
	chatHandler := chatHandler.NewChatHandler(chatClient, helper.NewHelper(&cfg))

	serverHTTP := server.NewServerHTTP(
		adminAuthHandler, jobseekerAuthHandler,
		recruiterAuthHandler, recruiterJobHandler,
		JobseekerJobhandler, jobseekerPostHandler,
		chatHandler, VideoHandler, authHandler, notificationHandler)

	return serverHTTP, nil
}
