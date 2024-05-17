package usecase

import (
	repo "ConnetHub_post/pkg/repository/interfaces"
	"ConnetHub_post/pkg/usecase/interfaces"
)

type recruiterJobUseCase struct {
	postRepository repo.RecruiterPostRepository
}

func NewRecruiterPostUseCase(repo repo.RecruiterPostRepository) interfaces.RecruiterPostUsecase {
	return &recruiterJobUseCase{
		postRepository: repo,
	}
}
