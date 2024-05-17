package di

import (
	"ConnetHub_video/pkg/api/server"
	"ConnetHub_video/pkg/config"
	"ConnetHub_video/pkg/db"
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
