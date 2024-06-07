package server

import (
	authHandler "connectHub_gateway/pkg/api/handler/auth"
	chatHandler "connectHub_gateway/pkg/api/handler/chat"
	jobHandler "connectHub_gateway/pkg/api/handler/job"
	postHandler "connectHub_gateway/pkg/api/handler/post"
	"connectHub_gateway/pkg/config"
	"connectHub_gateway/pkg/middleware"
	"log"

	"github.com/gin-gonic/gin"
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
	chatHandler *chatHandler.ChatHandler,

) *ServerHTTP {

	// Gin Engine
	router := gin.New()
	router.Use(gin.Logger())

	// Router Group
	adminRoute := router.Group("/admin")
	jobseekerRoute := router.Group("/jobseeker")
	recruiterRoute := router.Group("/recruiter")

	// Admin Routes
	adminRoute.POST("/login", AdminHandler.AdminLogin)

	// Jobseeker Routes
	jobseekerRoute.POST("/signup", JobseekerHandler.JobSeekerSignup)
	jobseekerRoute.POST("/login", JobseekerHandler.JobSeekerLogin)

	// Recruiter Routes
	recruiterRoute.POST("/signup", RecruiterHandler.RecruiterSignup)
	recruiterRoute.POST("/login", RecruiterHandler.RecruiterLogin)

	chat := router.Group("/user/chat")
	{
		chat.GET("", chatHandler.SendMessage)
		// chat.GET("/message", chatHandler.GetChat)
	}

	router.Use(middleware.AuthMiddleware)
	{

		// Admin Router Group
		adminJobseeker := router.Group("/admin/jobseeker")
		adminRecruiter := router.Group("/admin/recruiter")
		adminAuthRoute := router.Group("/admin")

		// Jobseeker Router Group
		jobseekerAuthRoute := router.Group("/jobseeker")

		// recruiter Router Group
		recruiterAuthRoute := router.Group("/recruiter")

		// Admin Routes
		adminAuthRoute.POST("/policy", AdminHandler.CreatePolicy)
		adminAuthRoute.PUT("/policy", AdminHandler.UpdatePolicy)
		adminAuthRoute.DELETE("/policy", AdminHandler.DeletePolicy)
		adminAuthRoute.GET("/policy", AdminHandler.GetOnePolicy)
		adminAuthRoute.GET("/policies", AdminHandler.GetAllPolicies)

		// Admin Routes jobseeker
		adminJobseeker.GET("/all", AdminHandler.GetJobseekers)
		adminJobseeker.PATCH("/block", AdminHandler.BlockJobseeker)
		adminJobseeker.PATCH("/unblock", AdminHandler.UnBlockJobseeker)
		adminJobseeker.GET("", AdminHandler.GetJobseekerDetails)

		// Admin Routes recruiter
		adminRecruiter.GET("/all", AdminHandler.GetRecruiters)
		adminRecruiter.PATCH("/block", AdminHandler.BlockRecruiter)
		adminRecruiter.PATCH("/unblock", AdminHandler.UnBlockRecruiter)
		adminRecruiter.GET("", AdminHandler.GetRecruiterDetails)

		// Jobseeker Routes
		jobseekerAuthRoute.GET("/profile", JobseekerHandler.JobSeekerGetProfile)
		jobseekerAuthRoute.PATCH("/profile", JobseekerHandler.JobSeekerEditProfile)

		jobseekerAuthRoute.GET("/jobs", JobseekerJobhandler.JobSeekerGetAllJobs)
		jobseekerAuthRoute.GET("/job", JobseekerJobhandler.JobSeekerGetJobByID)
		jobseekerAuthRoute.POST("/job", JobseekerJobhandler.JobSeekerApplyJob)
		jobseekerAuthRoute.GET("/appliedjobs", JobseekerJobhandler.GetAppliedJobs)

		jobseekerAuthRoute.POST("/post", jobseekerPostHandler.CreatePost)
		jobseekerAuthRoute.PATCH("/post", jobseekerPostHandler.UpdatePost)
		jobseekerAuthRoute.DELETE("/post", jobseekerPostHandler.DeletePost)
		jobseekerAuthRoute.GET("/post", jobseekerPostHandler.GetOnePost)
		jobseekerAuthRoute.GET("/posts", jobseekerPostHandler.GetAllPost)
		jobseekerAuthRoute.POST("/post/comment", jobseekerPostHandler.CreateCommentPost)
		jobseekerAuthRoute.PUT("/post/comment", jobseekerPostHandler.UpdateCommentPost)
		jobseekerAuthRoute.DELETE("/post/comment", jobseekerPostHandler.DeleteCommentPost)
		jobseekerAuthRoute.POST("/post/like", jobseekerPostHandler.AddLikePost)
		jobseekerAuthRoute.DELETE("/post/like", jobseekerPostHandler.RemoveLikePost)

		jobseekerAuthRoute.GET("/policy", JobseekerHandler.GetOnePolicy)
		jobseekerAuthRoute.GET("/policies", JobseekerHandler.GetAllPolicies)

		// Recruiter Routes
		recruiterAuthRoute.GET("/profile", RecruiterHandler.RecruiterGetProfile)
		recruiterAuthRoute.PATCH("/profile", RecruiterHandler.RecruiterEditProfile)

		recruiterAuthRoute.POST("/job", RecruiterJobHandler.PostJob)
		recruiterAuthRoute.GET("/jobs", RecruiterJobHandler.GetAllJobs)
		recruiterAuthRoute.GET("/job", RecruiterJobHandler.GetOneJob)
		recruiterAuthRoute.PATCH("/job", RecruiterJobHandler.UpdateAJob)
		recruiterAuthRoute.DELETE("/job", RecruiterJobHandler.DeleteAJob)
		recruiterAuthRoute.GET("/appliedjobs", RecruiterJobHandler.GetJobAppliedCandidates)
		recruiterAuthRoute.POST("/appliedjob", RecruiterJobHandler.ScheduleInterview)

		recruiterAuthRoute.GET("/policy", RecruiterHandler.GetOnePolicy)
		recruiterAuthRoute.GET("/policies", RecruiterHandler.GetAllPolicies)

		// chat
		chat := router.Group("/user/chat")
		{

			chat.GET("/message", chatHandler.GetChat)
		}

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
