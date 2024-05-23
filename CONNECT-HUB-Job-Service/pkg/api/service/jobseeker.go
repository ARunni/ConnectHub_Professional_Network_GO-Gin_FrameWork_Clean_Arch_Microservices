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

func (js *JobseekerJobServer) JobSeekerGetJobByID(ctx context.Context, req *jobseekerpb.JobSeekerGetJobByIDRequest) (*jobseekerpb.JobSeekerGetJobByIDResponse, error) {
	jobId := req.Id

	jobSeekerJobs, err := js.jobUseCase.JobSeekerGetJobByID(int(jobId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get job for job seeker: %v", err)
	}

	response := &jobseekerpb.JobSeekerGetJobByIDResponse{
		Job: &jobseekerpb.Job{
			Id:             uint64(jobSeekerJobs.ID),
			Title:          jobSeekerJobs.Title,
			Description:    jobSeekerJobs.Description,
			Location:       jobSeekerJobs.Location,
			EmployerId:     int64(jobSeekerJobs.EmployerID),
			EmploymentType: jobSeekerJobs.EmploymentType,
		},
	}

	return response, nil
}

func (js *JobseekerJobServer) JobSeekerApplyJob(ctx context.Context, req *jobseekerpb.JobSeekerApplyJobRequest) (*jobseekerpb.JobSeekerApplyJobResponse, error) {
	jobId := req.JobId
	jobseekerId := req.UserId

	jobSeekerJobs, err := js.jobUseCase.JobSeekerApplyJob(int(jobId),int(jobseekerId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to apply job for job seeker: %v", err)
	}

	response := &jobseekerpb.JobSeekerApplyJobResponse{
		Success: jobSeekerJobs,
	}

	return response, nil
}
