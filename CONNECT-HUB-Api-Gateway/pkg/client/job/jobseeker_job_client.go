package client

import (
	"os"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/job/interface"
	"github.com/ARunni/connectHub_gateway/pkg/config"

	"context"
	"fmt"

	jobseekerPb "github.com/ARunni/connectHub_gateway/pkg/pb/job/jobseeker"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type jobseekerJobClient struct {
	Client  jobseekerPb.JobseekerJobClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewJobseekerJobClient(cfg config.Config) interfaces.JobseekerJobClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	grpcConnection, err := grpc.Dial(cfg.ConnetHubJob, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth", err)
	}
	grpcClient := jobseekerPb.NewJobseekerJobClient(grpcConnection)
	return &jobseekerJobClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}

}

func (jc *jobseekerJobClient) JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error) {
	jc.Logger.Info("JobSeekerGetAllJobs at client started")
	resp, err := jc.Client.JobSeekerGetAllJobs(context.Background(), &jobseekerPb.JobSeekerGetAllJobsRequest{
		Title: keyword,
	})
	if err != nil {
		jc.Logger.Error("Error in JobSeekerGetAllJobs at client: ", err)
		return nil, fmt.Errorf("failed to get job: %v", err)
	}

	var jobs []models.JobSeekerGetAllJobs
	for _, job := range resp.Jobs {
		jobs = append(jobs, models.JobSeekerGetAllJobs{
			ID:    uint(job.Id),
			Title: job.Title,
		})
	}
	jc.Logger.Info("JobSeekerGetAllJobs at client success")

	return jobs, nil
}

func (jc *jobseekerJobClient) JobSeekerGetJobByID(id int) (models.JobOpeningData, error) {
	jc.Logger.Info("JobSeekerGetJobByID at client started")
	var job models.JobOpeningData
	resp, err := jc.Client.JobSeekerGetJobByID(context.Background(), &jobseekerPb.JobSeekerGetJobByIDRequest{
		Id: uint64(id),
	})
	if err != nil {
		jc.Logger.Error("Error in JobSeekerGetJobByID at client: ", err)
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

	jc.Logger.Info("JobSeekerGetJobByID at client success")

	return job, nil
}

func (jc *jobseekerJobClient) JobSeekerApplyJob(data models.ApplyJobReq) (models.ApplyJob, error) {
	jc.Logger.Info("JobSeekerApplyJob at client started")

	resp, err := jc.Client.JobSeekerApplyJob(context.Background(), &jobseekerPb.JobSeekerApplyJobRequest{
		JobId:       uint64(data.JobID),
		UserId:      uint64(data.JobseekerID),
		CoverLetter: data.CoverLetter,
		Resume:      data.Resume,
	})
	if err != nil {
		jc.Logger.Error("Error in JobSeekerApplyJob at client: ", err)
		return models.ApplyJob{}, fmt.Errorf("failed to apply job: %v", err)
	}

	jc.Logger.Info("JobSeekerApplyJob at client success")


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
	jc.Logger.Info("GetAppliedJobs at client started")


	resp, err := jc.Client.GetAppliedJobs(context.Background(), &jobseekerPb.JobSeekerGetAppliedJobsRequest{
		UserId: int64(user_id),
	})
	if err != nil {
		jc.Logger.Error("Error in GetAppliedJobs at client: ", err)
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
	jc.Logger.Info("GetAppliedJobs at client success")
	return models.AppliedJobsJ{Jobs: jobs}, nil
}
