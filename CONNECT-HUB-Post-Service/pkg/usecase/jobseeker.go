package usecase

import (
	repo "ConnetHub_post/pkg/repository/interfaces"
	"ConnetHub_post/pkg/usecase/interfaces"
)

type jobseekerJobUseCase struct {
	postRepository repo.JobseekerPostRepository
}

func NewJobseekerpostUseCase(repo repo.AdminPostRepository) interfaces.JobseekerPostUsecase {
	return &jobseekerJobUseCase{
		postRepository: repo,
	}
}
