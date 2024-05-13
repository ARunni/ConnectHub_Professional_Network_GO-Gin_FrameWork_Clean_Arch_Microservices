package service

import (
	"ConnetHub_job/pkg/pb/job"
	interfaces "ConnetHub_job/pkg/usecase/interface"
)

type JobServer struct {
	jobUseCase interfaces.JobUsecase
	job.UnimplementedJobServer
}

func NewJobServer(useCase interfaces.JobUsecase) job.JobServer {

	return &JobServer{
		jobUseCase: useCase,
	}
}
