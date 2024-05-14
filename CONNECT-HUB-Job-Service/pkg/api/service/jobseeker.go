package service

import (
	jobseekerpb "ConnetHub_job/pkg/pb/job/jobseeker"
	interfaces "ConnetHub_job/pkg/usecase/interface"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JobseekerJobServer struct {
	jobUseCase interfaces.JobseekerJobUsecase
	jobseekerpb.UnimplementedJobseekerJobServer
}

func NewJobseekerJobServer(useCase interfaces.JobseekerJobUsecase) jobseekerpb.JobseekerJobServer {

	return &JobseekerJobServer{
		jobUseCase: useCase,
	}
}

func (js *JobseekerJobServer) JobSeekerGetAllJobs(ctx context.Context, req *jobseekerpb.JobSeekerGetAllJobsRequest) (*jobseekerpb.JobSeekerGetAllJobsResponse, error) {
	keyword := req.Title

	jobSeekerJobs, err := js.jobUseCase.JobSeekerGetAllJobs(keyword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get jobs for job seeker: %v", err)
	}

	var jobsResponse []*jobseekerpb.JSGetAllJobsRespons
	for _, job := range jobSeekerJobs {
		jobResponse := &jobseekerpb.JSGetAllJobsRespons{
			Id:    uint64(job.ID),
			Title: job.Title,
		}
		jobsResponse = append(jobsResponse, jobResponse)
	}

	response := &jobseekerpb.JobSeekerGetAllJobsResponse{
		Jobs: jobsResponse,
	}

	return response, nil
}
