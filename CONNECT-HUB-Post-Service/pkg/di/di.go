package di

import (
	"ConnetHub_post/pkg/api/server"
	"ConnetHub_post/pkg/api/service"
	"ConnetHub_post/pkg/config"
	"ConnetHub_post/pkg/db"
	repo "ConnetHub_post/pkg/repository"
	usecase "ConnetHub_post/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	// recruiterJobRepository := repo.NewRecruiterJobRepository(gormDB)
	// recruiterJobUseCase := usecase.NewRecruiterJobUseCase(recruiterJobRepository)
	// recruiterJobServiceServer := service.NewRecruiterJobServer(recruiterJobUseCase)

	jobseekerPostRepository := repo.NewJobseekerPostRepository(gormDB)
	jobseekerPostUseCase := usecase.NewJobseekerpostUseCase(jobseekerPostRepository)
	jobseekerPostServiceServer := service.NewJobseekerPostServer(jobseekerPostUseCase)

	grpcServer, err := server.NewGRPCServer(
		cfg,
		jobseekerPostServiceServer,
		// recruiterJobServiceServer, jobseekerJobServiceServer
	)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
