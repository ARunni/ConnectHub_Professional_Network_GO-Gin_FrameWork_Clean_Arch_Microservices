package usecase

import (
	repo "ConnetHub_post/pkg/repository/interfaces"
	"ConnetHub_post/pkg/usecase/interfaces"
)

type adminJobUseCase struct {
	postRepository repo.AdminPostRepository
}

func NewAdminPostUseCase(repo repo.AdminPostRepository) interfaces.AdminPostUsecase {
	return &adminJobUseCase{
		postRepository: repo,
	}
}
