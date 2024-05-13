package di

import (
	"ConnetHub_job/pkg/api/server"
	service "ConnetHub_job/pkg/api/service"
	"ConnetHub_job/pkg/config"
	"ConnetHub_job/pkg/db"
	repo "ConnetHub_job/pkg/repository"
	usecase "ConnetHub_job/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	jobRepository := repo.NewJobRepository(gormDB)
	jobUseCase := usecase.NewJobUseCase(jobRepository)
	jobServiceServer := service.NewJobServer(jobUseCase)

	grpcServer, err := server.NewGRPCServer(
		cfg,
		jobServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
