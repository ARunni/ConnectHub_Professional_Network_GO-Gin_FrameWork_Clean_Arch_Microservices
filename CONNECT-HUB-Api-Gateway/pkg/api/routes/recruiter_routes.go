package routes

import (
	videoHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/Video_Call"
	authHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/auth"
	"github.com/ARunni/connectHub_gateway/pkg/middleware"
	"github.com/gin-gonic/gin"

	// AuthHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/auth"
	chatHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/chat"
	jobHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/job"
	notificationHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/notification"
)

func RecruiterRoutes(engine *gin.RouterGroup,
	authhandler *authHandler.RecruiterHandler,
	jobHandler *jobHandler.RecruiterJobHandler,
	videocallHandler *videoHandler.VideoCallHandler,
	notificationHandler *notificationHandler.NotificationHandler,
	chatHandler *chatHandler.ChatHandler, AuthHandlerVideo *authHandler.AuthHandler,
) {
	engine.POST("/signup", authhandler.RecruiterSignup)
	engine.POST("/login", authhandler.RecruiterLogin)

	engine.Use(middleware.RecruiterAuthMiddleware)
	{
		engine.GET("/profile", authhandler.RecruiterGetProfile)
		engine.PATCH("/profile", authhandler.RecruiterEditProfile)

		engine.POST("/job", jobHandler.PostJob)
		engine.GET("/jobs", jobHandler.GetAllJobs)
		engine.GET("/job", jobHandler.GetOneJob)
		engine.PATCH("/job", jobHandler.UpdateAJob)
		engine.DELETE("/job", jobHandler.DeleteAJob)
		engine.GET("/appliedjobs", jobHandler.GetJobAppliedCandidates)
		engine.POST("/appliedjob", jobHandler.ScheduleInterview)
		engine.DELETE("/interview", jobHandler.CancelScheduledInterview)

		engine.GET("/policy", authhandler.GetOnePolicy)
		engine.GET("/policies", authhandler.GetAllPolicies)

		engine.GET("chat/message", chatHandler.GetChat)

		engine.GET("/key", AuthHandlerVideo.VideoCallKey)

		notification := engine.Group("/notifications")
		{
			notification.GET("", notificationHandler.GetNotification)
			notification.PATCH("", notificationHandler.ReadNotification)
			notification.PATCH("/all", notificationHandler.MarkAllAsRead)
			notification.GET("/all", notificationHandler.GetAllNotifications)
		}

	}

}
