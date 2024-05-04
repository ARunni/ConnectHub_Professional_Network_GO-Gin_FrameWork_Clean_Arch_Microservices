package server

import (
	"connectHub_gateway/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP() *ServerHTTP {
	router := gin.New()
	router.Use(gin.Logger())

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
