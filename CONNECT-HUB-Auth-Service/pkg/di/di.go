package di

import (
	"github.com/ARunni/ConnetHub_auth/pkg/api/server"
	"github.com/ARunni/ConnetHub_auth/pkg/api/service"
	"github.com/ARunni/ConnetHub_auth/pkg/config"
	"github.com/ARunni/ConnetHub_auth/pkg/db"
	"github.com/ARunni/ConnetHub_auth/pkg/repository"
	"github.com/ARunni/ConnetHub_auth/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminServiceServer := service.NewAdminServer(adminUseCase)

	jobseekerRepository := repository.NewJobseekerRepository(gormDB)
	jobseekerUseCase := usecase.NewJobseekerUseCase(jobseekerRepository)
	jobseekerServiceServer := service.NewJobSeekerServer(jobseekerUseCase)

	recruiterRepository := repository.NewRecruiterRepository(gormDB)
	recruiterUseCase := usecase.NewRecruiterUseCase(recruiterRepository)
	recruiterServiceServer := service.NewRecruiterServer(recruiterUseCase)

	authVideoRepository := repository.NewVideoCallRepository(gormDB)
	authVideoUseCase := usecase.NewVideoCallUseCase(authVideoRepository)
	authVideoServiceServer := service.NewauthServer(authVideoUseCase)

	jobAuthServiceServer := service.NewJobauthServer(recruiterUseCase)

	grpcServer, err := server.NewGRPCServer(
		cfg,
		adminServiceServer,
		recruiterServiceServer,
		jobseekerServiceServer,
		jobAuthServiceServer,
		authVideoServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
