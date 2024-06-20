package di

import (
	"github.com/ARunni/ConnetHub_post/pkg/api/server"
	"github.com/ARunni/ConnetHub_post/pkg/api/service"
	auth "github.com/ARunni/ConnetHub_post/pkg/client/auth"
	"github.com/ARunni/ConnetHub_post/pkg/config"
	"github.com/ARunni/ConnetHub_post/pkg/db"
	repo "github.com/ARunni/ConnetHub_post/pkg/repository"
	usecase "github.com/ARunni/ConnetHub_post/pkg/usecase"
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
	authClient := auth.NewAuthClient(&cfg)
	jobseekerPostUseCase := usecase.NewJobseekerpostUseCase(jobseekerPostRepository, authClient)
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
