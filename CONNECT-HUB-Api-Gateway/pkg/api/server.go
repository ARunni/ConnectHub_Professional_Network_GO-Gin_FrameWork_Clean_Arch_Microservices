package server

import (
	authHandler "connectHub_gateway/pkg/api/handler/auth"
	jobHandler "connectHub_gateway/pkg/api/handler/job"
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

	router.Use(middleware.AuthMiddleware)
	{

		// Admin Router Group
		adminJobseeker := router.Group("/admin/jobseeker")
		adminRecruiter := router.Group("/admin/recruiter")

		// Jobseeker Router Group
		jobseekerAuthRoute := router.Group("/jobseeker")

		// recruiter Router Group
		recruiterAuthRoute := router.Group("/recruiter")

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

		// Recruiter Routes
		recruiterAuthRoute.GET("/profile", RecruiterHandler.RecruiterGetProfile)
		recruiterAuthRoute.PATCH("/profile", RecruiterHandler.RecruiterEditProfile)
		recruiterAuthRoute.POST("/job", RecruiterJobHandler.PostJob)
		recruiterAuthRoute.GET("/jobs", RecruiterJobHandler.GetAllJobs)
		recruiterAuthRoute.GET("/job", RecruiterJobHandler.GetOneJob)
		recruiterAuthRoute.PATCH("/job", RecruiterJobHandler.UpdateAJob)
		recruiterAuthRoute.DELETE("/job", RecruiterJobHandler.DeleteAJob)
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
