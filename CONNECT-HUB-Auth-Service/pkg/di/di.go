package di

import (
	"ConnetHub_auth/pkg/api/server"
	"ConnetHub_auth/pkg/api/service"
	"ConnetHub_auth/pkg/config"
	"ConnetHub_auth/pkg/db"
	"ConnetHub_auth/pkg/repository"
	"ConnetHub_auth/pkg/usecase"
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

	jobAuthServiceServer := service.NewJobauthServer(recruiterUseCase)

	grpcServer, err := server.NewGRPCServer(
		cfg,
		adminServiceServer,
		recruiterServiceServer,
		jobseekerServiceServer,
		jobAuthServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
