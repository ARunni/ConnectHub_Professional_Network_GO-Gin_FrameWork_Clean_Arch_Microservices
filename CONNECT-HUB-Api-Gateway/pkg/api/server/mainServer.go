package server

import (
	"log"

	videoHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/Video_Call"
	authHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/auth"
	"github.com/ARunni/connectHub_gateway/pkg/api/routes"
	"github.com/ARunni/connectHub_gateway/pkg/config"

	// AuthHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/auth"
	chatHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/chat"
	jobHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/job"
	notificationHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/notification"
	postHandler "github.com/ARunni/connectHub_gateway/pkg/api/handler/post"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	AdminHandler *authHandler.AdminHandler,
	JobseekerHandler *authHandler.JobSeekerHandler,
	RecruiterHandler *authHandler.RecruiterHandler,
	RecruiterJobHandler *jobHandler.RecruiterJobHandler,
	JobseekerJobhandler *jobHandler.JobseekerJobHandler,
	jobseekerPostHandler *postHandler.JobseekerPostHandler,
	chatHandler *chatHandler.ChatHandler, videocallHandler *videoHandler.VideoCallHandler,
	AuthHandler *authHandler.AuthHandler, notificationHandler *notificationHandler.NotificationHandler,

) *ServerHTTP {
	// Gin Engine
	router := gin.New()
	router.Use(gin.Logger())

	// router.LoadHTMLGlob("pkg/templates/index.html")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/static", "./static")
	router.LoadHTMLGlob("template/*")

	router.GET("/exit", videocallHandler.ExitPage)
	router.GET("/error", videocallHandler.ErrorPage)
	router.GET("/index", videocallHandler.IndexedPage)

	routes.AdminRoutes(router.Group("/admin"), AdminHandler)

	routes.JobseekerRoutes(router.Group("/jobseeker"),
		JobseekerHandler, JobseekerJobhandler,
		videocallHandler, notificationHandler,
		chatHandler, jobseekerPostHandler)

	routes.RecruiterRoutes(router.Group("/recruiter"),
		RecruiterHandler, RecruiterJobHandler,
		videocallHandler, notificationHandler,
		chatHandler, AuthHandler)

	chat := router.Group("/user/chat")
	{
		chat.GET("", chatHandler.SendMessage)
	}

	return &ServerHTTP{engine: router}

}

func (s *ServerHTTP) Start() {

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Printf("error while loading the config")
	}

	log.Printf("starting server on :7000")

	err = s.engine.Run(cfg.Port)

	if err != nil {
		log.Printf("error while starting the server")
	}
}
