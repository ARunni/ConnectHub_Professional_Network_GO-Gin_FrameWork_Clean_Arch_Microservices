package service

import "ConnetHub_post/pkg/usecase/interfaces"

type RecruiterPostServer struct {
	postUseCase interfaces.RecruiterPostUsecase
	jobpb.UnimplementedRecruiterJobServer
}

func NewRecruiterPostServer(useCase interfaces.RecruiterPostUsecase) jobpb.RecruiterJobServer {

	return &RecruiterPostServer{
		postUseCase: useCase,
	}
}
