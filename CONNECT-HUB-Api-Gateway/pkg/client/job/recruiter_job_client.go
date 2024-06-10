package client

import (
	logging "connectHub_gateway/Logging"
	interfaces "connectHub_gateway/pkg/client/job/interface"
	"connectHub_gateway/pkg/config"
	recruiterPb "connectHub_gateway/pkg/pb/job/recruiter"
	"connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type recruiterJobClient struct {
	Client  recruiterPb.RecruiterJobClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewRecruiterJobClient(cfg config.Config) interfaces.RecruiterJobClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	grpcConnection, err := grpc.Dial(cfg.ConnetHubJob, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth", err)
	}
	grpcClient := recruiterPb.NewRecruiterJobClient(grpcConnection)
	return &recruiterJobClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}

}
func (jc *recruiterJobClient) PostJob(data models.JobOpening) (models.JobOpeningData, error) {
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

func (jc *recruiterJobClient) GetAllJobs(recruiterID int32) ([]models.AllJob, error) {

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

func (jc *recruiterJobClient) GetOneJob(recruiterID, jobId int32) (models.JobOpeningData, error) {

	resp, err := jc.Client.GetOneJob(context.Background(), &recruiterPb.GetAJobRequest{
		EmployerIDInt: recruiterID,
		JobId:         jobId,
	})
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to get job: %v", err)
	}

	postedOnTime := resp.PostedOn.AsTime()
	applicationDeadlineTime := resp.ApplicationDeadline.AsTime()
	salary, _ := strconv.Atoi(resp.Salary)
	// if err != nil {
	// 	return models.JobOpeningData{}, fmt.Errorf("failed to convert salary to int: %v", err)
	// }
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

func (jc *recruiterJobClient) DeleteAJob(employerIDInt, jobID int32) error {
	_, err := jc.Client.DeleteAJob(context.Background(), &recruiterPb.DeleteAJobRequest{EmployerIDInt: employerIDInt, JobId: jobID})
	if err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}
	return nil
}

func (jc *recruiterJobClient) UpdateAJob(employerIDInt int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningData, error) {

	applicationDeadline := timestamppb.New(jobDetails.ApplicationDeadline)

	job, err := jc.Client.UpdateAJob(context.Background(), &recruiterPb.UpdateAJobRequest{
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              jobDetails.Salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: applicationDeadline,
		EmployerId:          employerIDInt,
		JobId:               jobID,
	})
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to post job opening: %v", err)
	}

	postedOnTime := job.PostedOn.AsTime()
	applicationDeadlineTime := job.ApplicationDeadline.AsTime()
	salary, _ := strconv.Atoi(job.Salary)
	return models.JobOpeningData{
		ID:                  uint(job.Id),
		Title:               job.Title,
		Description:         job.Description,
		Requirements:        job.Requirements,
		PostedOn:            postedOnTime,
		Location:            job.Location,
		EmploymentType:      job.EmploymentType,
		Salary:              salary,
		SkillsRequired:      job.SkillsRequired,
		ExperienceLevel:     job.ExperienceLevel,
		EducationLevel:      job.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
		EmployerID:          int(job.EmployerId),
	}, nil

}

func (jc *recruiterJobClient) GetJobAppliedCandidates(recruiter_id int) (models.AppliedJobs, error) {

	job, err := jc.Client.GetJobAppliedCandidates(context.Background(), &recruiterPb.GetAppliedJobsRequest{
		UserId: int64(recruiter_id),
	})
	if err != nil {
		return models.AppliedJobs{}, fmt.Errorf("failed to apply job: %v", err)
	}
	var jobs []models.ApplyJobRes
	for _, job := range job.Jobs {
		jobs = append(jobs, models.ApplyJobRes{
			ID:            uint(job.Id),
			JobID:         uint(job.JobId),
			JobseekerID:   uint(job.UserId),
			RecruiterID:   uint(job.RecruiterId),
			Status:        job.Status,
			CoverLetter:   job.CoverLetter,
			ResumeUrl:     job.ResumeUrl,
			JobseekerName: job.JobseekerName,
			JoseekerEmail: job.JobseekerEmail,
		})
		fmt.Println("hjhjhjhjhjjhhjk", job.JobseekerEmail)
	}
	return models.AppliedJobs{Jobs: jobs}, nil
}

func (jc *recruiterJobClient) ScheduleInterview(data models.ScheduleReq) (models.Interview, error) {

	job, err := jc.Client.ScheduleInterview(context.Background(), &recruiterPb.ScheduleInterviewRequest{
		ApplicationId: int64(data.ApplicationId),
		RecruiterId:   int64(data.RecruiterID),
		DateAndTime:   timestamppb.New(data.DateAndTime),
		Mode:          data.Mode,
		Link:          data.Link,
	})
	if err != nil {
		return models.Interview{}, fmt.Errorf("failed to apply job: %v", err)
	}

	return models.Interview{
		ID:            uint(job.Id),
		JobID:         uint(job.JobId),
		JobseekerID:   uint(job.JobseekerId),
		RecruiterID:   uint(job.RecruiterId),
		DateAndTime:   job.DateAndTime.AsTime(),
		Mode:          job.Mode,
		Link:          job.Link,
		Status:        job.Status,
		ApplicationId: uint(job.ApplicationId),
	}, nil
}
