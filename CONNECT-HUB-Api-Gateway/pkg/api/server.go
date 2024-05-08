package server

import (
	"connectHub_gateway/pkg/api/handler"
	"connectHub_gateway/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	AdminHandler *handler.AdminHandler,
	JobseekerHandler *handler.JobSeekerHandler,
	RecruiterHandler *handler.RecruiterHandler,
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
