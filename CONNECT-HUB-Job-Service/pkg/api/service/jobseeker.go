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

	jobSeekerJobs, err := js.jobUseCase.JobSeekerApplyJob(int(jobId), int(jobseekerId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to apply job for job seeker: %v", err)
	}

	response := &jobseekerpb.JobSeekerApplyJobResponse{
		Success: jobSeekerJobs,
	}

	return response, nil
}

func (js *JobseekerJobServer) GetAppliedJobs(ctx context.Context, req *jobseekerpb.JobSeekerGetAppliedJobsRequest) (*jobseekerpb.JobSeekerGetAppliedJobsResponse, error) {

	jobseekerId := req.UserId

	jobSeekerJobs, err := js.jobUseCase.GetAppliedJobs(int(jobseekerId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get applied job for job seeker: %v", err)
	}
	var jobs []*jobseekerpb.AppliedJobs
	for _, job := range jobSeekerJobs.Jobs {
		jobs = append(jobs, &jobseekerpb.AppliedJobs{
			JobId:       int64(job.JobID),
			Id:          int64(job.ID),
			UserId:      int64(job.JobseekerID),
			RecruiterId: int64(job.RecruiterID),
			Status:      job.Status,
		})
	}
	response := &jobseekerpb.JobSeekerGetAppliedJobsResponse{
		Jobs: jobs,
	}

	return response, nil
}
