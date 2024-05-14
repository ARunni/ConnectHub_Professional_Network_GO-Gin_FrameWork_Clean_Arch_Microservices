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

	recruiterJobRepository := repo.NewRecruiterJobRepository(gormDB)
	recruiterJobUseCase := usecase.NewRecruiterJobUseCase(recruiterJobRepository)
	recruiterJobServiceServer := service.NewRecruiterJobServer(recruiterJobUseCase)

	jobseekerJobRepository := repo.NewjobseekerJobRepository(gormDB)
	jobseekerJobUseCase := usecase.NewJobseekerJobUseCase(jobseekerJobRepository)
	jobseekerJobServiceServer := service.NewJobseekerJobServer(jobseekerJobUseCase)

	grpcServer, err := server.NewGRPCServer(
		cfg,
		recruiterJobServiceServer, jobseekerJobServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
