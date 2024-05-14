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
