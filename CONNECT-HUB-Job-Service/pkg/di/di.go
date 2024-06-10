package di

import (
	"github.com/ARunni/ConnetHub_job/pkg/api/server"
	service "github.com/ARunni/ConnetHub_job/pkg/api/service"
	jobAuth "github.com/ARunni/ConnetHub_job/pkg/client/auth"
	"github.com/ARunni/ConnetHub_job/pkg/config"
	"github.com/ARunni/ConnetHub_job/pkg/db"
	repo "github.com/ARunni/ConnetHub_job/pkg/repository"
	usecase "github.com/ARunni/ConnetHub_job/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	recruiterJobRepository := repo.NewRecruiterJobRepository(gormDB)
	jobAuthClient := jobAuth.NewJobAuthClient(cfg)
	recruiterJobUseCase := usecase.NewRecruiterJobUseCase(recruiterJobRepository, jobAuthClient)
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
