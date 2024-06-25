package routes

import (
	videoHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/Video_Call"
	authhandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/auth"
	chathandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/chat"
	"github.com/ARunni/connectHub_gateway/pkg/middleware"
	"github.com/gin-gonic/gin"

	// AuthHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/auth"
	jobHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/job"
	notificationHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/notification"
	postHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/post"
)

func JobseekerRoutes(engine *gin.RouterGroup, authhandler *authhandler.JobSeekerHandler,
	jobHandler *jobHandler.JobseekerJobHandler, videocallHandler *videoHandler.VideoCallHandler,
	notificationHandler *notificationHandler.NotificationHandler, chatHandler *chathandler.ChatHandler,
	postHandler *postHandler.JobseekerPostHandler) {

	engine.POST("/signup", authhandler.JobSeekerSignup)
	engine.POST("/login", authhandler.JobSeekerLogin)

	engine.Use(middleware.JobseekerAuthMiddleware)
	{
		engine.GET("/profile", authhandler.JobSeekerGetProfile)
		engine.PATCH("/profile", authhandler.JobSeekerEditProfile)

		engine.GET("/jobs", jobHandler.JobSeekerGetAllJobs)
		engine.GET("/job", jobHandler.JobSeekerGetJobByID)
		engine.POST("/job", jobHandler.JobSeekerApplyJob)
		engine.GET("/appliedjobs", jobHandler.GetAppliedJobs)

		engine.POST("/post", postHandler.CreatePost)
		engine.PATCH("/post", postHandler.UpdatePost)
		engine.DELETE("/post", postHandler.DeletePost)
		engine.GET("/post", postHandler.GetOnePost)
		engine.GET("/posts", postHandler.GetAllPost)
		engine.POST("/post/comment", postHandler.CreateCommentPost)
		engine.PUT("/post/comment", postHandler.UpdateCommentPost)
		engine.DELETE("/post/comment", postHandler.DeleteCommentPost)
		engine.POST("/post/like", postHandler.AddLikePost)
		engine.DELETE("/post/like", postHandler.RemoveLikePost)

		engine.GET("/policy", authhandler.GetOnePolicy)
		engine.GET("/policies", authhandler.GetAllPolicies)

		engine.GET("chat/message", chatHandler.GetChatJobseeker)

		notification := engine.Group("/notifications")

		notification.GET("", notificationHandler.GetNotification)
		notification.PATCH("", notificationHandler.ReadNotification)
		notification.PATCH("/all", notificationHandler.MarkAllAsRead)
		notification.GET("/all", notificationHandler.GetAllNotifications)
	}

}
