package client

import (
	interfaces "connectHub_gateway/pkg/client/interface"
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
	Client recruiterPb.JobClient
}

func NewJobClient(cfg config.Config) interfaces.JobClient {

	grpcConnection, err := grpc.Dial(cfg.ConnetHubJob, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth", err)
	}
	grpcClient := recruiterPb.NewJobClient(grpcConnection)
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
