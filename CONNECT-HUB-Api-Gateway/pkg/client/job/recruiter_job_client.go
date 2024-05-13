package client

import (
	interfaces "connectHub_gateway/pkg/client/job/interface"
	"connectHub_gateway/pkg/config"
	recruiterPb "connectHub_gateway/pkg/pb/job/recruiter"
	"connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type jobClient struct {
	Client recruiterPb.RecruiterJobClient
}

func NewRecruiterJobClient(cfg config.Config) interfaces.RecruiterJobClient {

	grpcConnection, err := grpc.Dial(cfg.ConnetHubJob, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth", err)
	}
	grpcClient := recruiterPb.NewRecruiterJobClient(grpcConnection)
	return &jobClient{
		Client: grpcClient,
	}

}
func (jc *jobClient) PostJob(data models.JobOpening) (models.JobOpeningData, error) {
	applicationDeadline := timestamppb.New(data.ApplicationDeadline)
	job, err := jc.Client.PostJob(context.Background(), &recruiterPb.JobOpeningRequest{
		Title:               data.Title,
		Description:         data.Description,
		Requirements:        data.Requirements,
		EmployerId:          int32(data.EmployerID),
		Location:            data.Location,
		EmploymentType:      data.EmploymentType,
		Salary:              data.Salary,
		SkillsRequired:      data.SkillsRequired,
		ExperienceLevel:     data.ExperienceLevel,
		EducationLevel:      data.EducationLevel,
		ApplicationDeadline: applicationDeadline,
	})
	if err != nil {
		return models.JobOpeningData{}, err
	}
	postedOnTime := job.PostedOn.AsTime()
	applicationDeadlineTime := job.ApplicationDeadline.AsTime()

	salary, err := strconv.Atoi(job.Salary)
	if err != nil {
		return models.JobOpeningData{}, err
	}
	return models.JobOpeningData{
		ID:                  uint(job.Id),
		Title:               job.Title,
		Description:         job.Description,
		Requirements:        job.Requirements,
		PostedOn:            postedOnTime,
		EmployerID:          data.EmployerID,
		Location:            job.Location,
		EmploymentType:      job.EmploymentType,
		Salary:              salary,
		SkillsRequired:      job.SkillsRequired,
		ExperienceLevel:     job.ExperienceLevel,
		EducationLevel:      job.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
	}, nil
}

func (jc *jobClient) GetAllJobs(recruiterID int32) ([]models.AllJob, error) {

	resp, err := jc.Client.GetAllJobs(context.Background(), &recruiterPb.GetAllJobsRequest{EmployerIDInt: recruiterID})
	if err != nil {
		return nil, fmt.Errorf("failed to get all jobs: %v", err)
	}

	var allJobs []models.AllJob
	for _, job := range resp.Jobs {

		applicationDeadlineTime := job.ApplicationDeadline.AsTime()

		allJobs = append(allJobs, models.AllJob{
			ID:                  uint(job.Id),
			Title:               job.Title,
			ApplicationDeadline: applicationDeadlineTime,
			EmployerID:          recruiterID,
		})
	}

	return allJobs, nil
}

func (jc *jobClient) GetOneJob(recruiterID, jobId int32) (models.JobOpeningData, error) {

	resp, err := jc.Client.GetOneJob(context.Background(), &recruiterPb.GetAJobRequest{
		EmployerIDInt: recruiterID,
		JobId:         jobId,
	})
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to get job: %v", err)
	}

	postedOnTime := resp.PostedOn.AsTime()
	applicationDeadlineTime := resp.ApplicationDeadline.AsTime()
	salary, err := strconv.Atoi(resp.Salary)
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to convert salary to int: %v", err)
	}
	return models.JobOpeningData{
		ID:                  uint(resp.Id),
		Title:               resp.Title,
		Description:         resp.Description,
		Requirements:        resp.Requirements,
		PostedOn:            postedOnTime,
		Location:            resp.Location,
		EmploymentType:      resp.EmploymentType,
		Salary:              salary,
		SkillsRequired:      resp.SkillsRequired,
		ExperienceLevel:     resp.ExperienceLevel,
		EducationLevel:      resp.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
		EmployerID:          int(recruiterID),
	}, nil
}

func (jc *jobClient) DeleteAJob(employerIDInt, jobID int32) error {
	_, err := jc.Client.DeleteAJob(context.Background(), &recruiterPb.DeleteAJobRequest{EmployerIDInt: employerIDInt, JobId: jobID})
	if err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}
	return nil
}
