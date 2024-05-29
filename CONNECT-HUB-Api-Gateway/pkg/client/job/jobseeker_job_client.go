package client

import (
	interfaces "connectHub_gateway/pkg/client/job/interface"
	"connectHub_gateway/pkg/config"
	"connectHub_gateway/pkg/pb/job/jobseeker"
	jobseekerPb "connectHub_gateway/pkg/pb/job/jobseeker"
	"connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type jobseekerJobClient struct {
	Client jobseeker.JobseekerJobClient
}

func NewJobseekerJobClient(cfg config.Config) interfaces.JobseekerJobClient {

	grpcConnection, err := grpc.Dial(cfg.ConnetHubJob, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth", err)
	}
	grpcClient := jobseekerPb.NewJobseekerJobClient(grpcConnection)
	return &jobseekerJobClient{
		Client: grpcClient,
	}

}

func (jc *jobseekerJobClient) JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error) {
	resp, err := jc.Client.JobSeekerGetAllJobs(context.Background(), &jobseekerPb.JobSeekerGetAllJobsRequest{
		Title: keyword,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %v", err)
	}

	var jobs []models.JobSeekerGetAllJobs
	for _, job := range resp.Jobs {
		jobs = append(jobs, models.JobSeekerGetAllJobs{
			ID:    uint(job.Id),
			Title: job.Title,
		})
	}

	return jobs, nil
}

func (jc *jobseekerJobClient) JobSeekerGetJobByID(id int) (models.JobOpeningData, error) {
	var job models.JobOpeningData
	resp, err := jc.Client.JobSeekerGetJobByID(context.Background(), &jobseekerPb.JobSeekerGetJobByIDRequest{
		Id: uint64(id),
	})
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to get job: %v", err)
	}

	job.ID = uint(resp.Job.Id)
	job.Title = resp.Job.Title
	job.Description = resp.Job.Description
	job.Location = resp.Job.Location
	job.EmployerID = int(resp.Job.EmployerId)
	job.EmploymentType = resp.Job.EmploymentType

	return job, nil
}

func (jc *jobseekerJobClient) JobSeekerApplyJob(jobId, userId int) (bool, error) {

	resp, err := jc.Client.JobSeekerApplyJob(context.Background(), &jobseekerPb.JobSeekerApplyJobRequest{
		JobId:  uint64(jobId),
		UserId: uint64(userId),
	})
	if err != nil {
		return false, fmt.Errorf("failed to apply job: %v", err)
	}

	return resp.Success, nil
}

func (jc *jobseekerJobClient) GetAppliedJobs(user_id int) (models.AppliedJobs, error) {

	resp, err := jc.Client.GetAppliedJobs(context.Background(), &jobseekerPb.JobSeekerGetAppliedJobsRequest{
		UserId: int64(user_id),
	})
	if err != nil {
		return models.AppliedJobs{}, fmt.Errorf("failed to apply job: %v", err)
	}
	var jobs []models.ApplyJob
	for _, job := range resp.Jobs {
		jobs = append(jobs, models.ApplyJob{
			ID:          uint(job.Id),
			JobID:       uint(job.JobId),
			JobseekerID: uint(job.UserId),
			RecruiterID: uint(job.RecruiterId),
			Status:      job.Status,
		})
	}
	return models.AppliedJobs{Jobs: jobs}, nil
}
