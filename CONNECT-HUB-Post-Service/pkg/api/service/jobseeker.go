package service

import "ConnetHub_post/pkg/usecase/interfaces"

type JobseekerPostServer struct {
	postUseCase interfaces.JobseekerPostUsecase
	jobpb.UnimplementedRecruiterJobServer
}

func NewJobseekerPostServer(useCase interfaces.JobseekerPostUsecase) jobpb.RecruiterJobServer {

	return &JobseekerPostServer{
		postUseCase: useCase,
	}
}
