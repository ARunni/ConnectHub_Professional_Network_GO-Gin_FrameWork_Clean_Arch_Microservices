package client

import (
	interfaces "connectHub_gateway/pkg/client/job/interface"
	"connectHub_gateway/pkg/config"

	jobseekerPb "connectHub_gateway/pkg/pb/job/jobseeker"
	"connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type jobseekerJobClient struct {
	Client jobseekerPb.JobseekerJobClient
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
	job.Salary = int(resp.Job.Salary)
	job.SkillsRequired = resp.Job.SkillsRequired
	job.ExperienceLevel = resp.Job.ExperienceLevel
	job.EducationLevel = resp.Job.EducationLevel
	job.ApplicationDeadline = resp.Job.ApplicationDeadline.AsTime()

	return job, nil
}

func (jc *jobseekerJobClient) JobSeekerApplyJob(data models.ApplyJobReq) (models.ApplyJob, error) {

	resp, err := jc.Client.JobSeekerApplyJob(context.Background(), &jobseekerPb.JobSeekerApplyJobRequest{
		JobId:       uint64(data.JobID),
		UserId:      uint64(data.JobseekerID),
		CoverLetter: data.CoverLetter,
		Resume:      data.Resume,
	})
	if err != nil {
		return models.ApplyJob{}, fmt.Errorf("failed to apply job: %v", err)
	}

	return models.ApplyJob{
		ID:          uint(resp.Job.Id),
		JobID:       uint(resp.Job.JobId),
		JobseekerID: uint(resp.Job.UserId),
		RecruiterID: uint(resp.Job.RecruiterId),
		CoverLetter: resp.Job.CoverLetter,
		ResumeUrl:   resp.Job.ResumeUrl,
		Status:      resp.Job.Status,
	}, nil
}

func (jc *jobseekerJobClient) GetAppliedJobs(user_id int) (models.AppliedJobsJ, error) {

	resp, err := jc.Client.GetAppliedJobs(context.Background(), &jobseekerPb.JobSeekerGetAppliedJobsRequest{
		UserId: int64(user_id),
	})
	if err != nil {
		return models.AppliedJobsJ{}, fmt.Errorf("failed to apply job: %v", err)
	}
	var jobs []models.ApplyJob
	for _, job := range resp.Jobs {
		jobs = append(jobs, models.ApplyJob{
			ID:          uint(job.Id),
			JobID:       uint(job.JobId),
			JobseekerID: uint(job.UserId),
			RecruiterID: uint(job.RecruiterId),
			Status:      job.Status,
			CoverLetter: job.CoverLetter,
			ResumeUrl:   job.ResumeUrl,
		})
	}
	return models.AppliedJobsJ{Jobs: jobs}, nil
}
