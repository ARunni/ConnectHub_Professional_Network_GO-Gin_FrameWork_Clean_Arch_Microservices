package service

import (
	job "ConnetHub_job/pkg/pb/job/recruiter"
	interfaces "ConnetHub_job/pkg/usecase/interface"
	"ConnetHub_job/pkg/utils/models"
	"context"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type JobServer struct {
	jobUseCase interfaces.JobUsecase
	job.UnimplementedRecruiterJobServer
}

func NewJobServer(useCase interfaces.JobUsecase) job.RecruiterJobServer {

	return &JobServer{
		jobUseCase: useCase,
	}
}

func (js *JobServer) PostJob(ctx context.Context, Req *job.JobOpeningRequest) (*job.JobOpeningResponse, error) {
	applicationDeadlineTime := Req.ApplicationDeadline.AsTime()

	recruiterJob := models.JobOpening{
		EmployerID:          int(Req.EmployerId),
		Title:               Req.Title,
		Description:         Req.Description,
		Requirements:        Req.Requirements,
		Location:            Req.Location,
		EmploymentType:      Req.EmploymentType,
		Salary:              Req.Salary,
		SkillsRequired:      Req.SkillsRequired,
		ExperienceLevel:     Req.ExperienceLevel,
		EducationLevel:      Req.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
	}
	jobData, err := js.jobUseCase.PostJob(recruiterJob)
	if err != nil {
		return nil, err
	}
	salary := strconv.Itoa(jobData.Salary)
	postedOnTime := timestamppb.New(jobData.PostedOn)
	return &job.JobOpeningResponse{
		Id:                  uint64(jobData.ID),
		Title:               jobData.Title,
		Description:         jobData.Description,
		Requirements:        jobData.Requirements,
		PostedOn:            postedOnTime,
		EmployerId:          int32(jobData.EmployerID),
		Location:            jobData.Location,
		EmploymentType:      jobData.EmploymentType,
		Salary:              salary,
		SkillsRequired:      jobData.SkillsRequired,
		ExperienceLevel:     jobData.ExperienceLevel,
		EducationLevel:      jobData.EducationLevel,
		ApplicationDeadline: timestamppb.New(jobData.ApplicationDeadline),
	}, nil
}
