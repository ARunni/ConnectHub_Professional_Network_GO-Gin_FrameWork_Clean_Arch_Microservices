package usecase

import (
	repo "ConnetHub_job/pkg/repository/interface"
	interfaces "ConnetHub_job/pkg/usecase/interface"
)

type jobUseCase struct {
	jobRepository repo.JobRepository
}

func NewJobUseCase(repo repo.JobRepository) interfaces.JobUsecase {
	return &jobUseCase{
		jobRepository: repo,
	}
}
