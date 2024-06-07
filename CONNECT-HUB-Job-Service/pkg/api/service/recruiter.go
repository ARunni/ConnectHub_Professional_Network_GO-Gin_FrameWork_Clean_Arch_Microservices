package service

import (
	jobpb "ConnetHub_job/pkg/pb/job/recruiter"
	interfaces "ConnetHub_job/pkg/usecase/interface"
	"ConnetHub_job/pkg/utils/models"
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RecruiterJobServer struct {
	jobUseCase interfaces.RecruiterJobUsecase
	jobpb.UnimplementedRecruiterJobServer
}

func NewRecruiterJobServer(useCase interfaces.RecruiterJobUsecase) jobpb.RecruiterJobServer {

	return &RecruiterJobServer{
		jobUseCase: useCase,
	}
}

func (js *RecruiterJobServer) PostJob(ctx context.Context, Req *jobpb.JobOpeningRequest) (*jobpb.JobOpeningResponse, error) {
	applicationDeadlineTime := Req.ApplicationDeadline.AsTime()
	fmt.Println("refdssd id ", Req.EmployerId)
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
	return &jobpb.JobOpeningResponse{
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

func (js *RecruiterJobServer) GetAllJobs(ctx context.Context, req *jobpb.GetAllJobsRequest) (*jobpb.GetAllJobsResponse, error) {
	employerID := int32(req.EmployerIDInt)

	jobs, err := js.jobUseCase.GetAllJobs(employerID)
	if err != nil {
		return nil, err
	}

	var jobResponses []*jobpb.JobOpeningResponse
	for _, job := range jobs {
		jobResponse := &jobpb.JobOpeningResponse{
			Id:                  uint64(job.ID),
			Title:               job.Title,
			ApplicationDeadline: timestamppb.New(job.ApplicationDeadline),
			EmployerId:          job.EmployerID,
		}
		jobResponses = append(jobResponses, jobResponse)
	}

	return &jobpb.GetAllJobsResponse{Jobs: jobResponses}, nil
}

func (js *RecruiterJobServer) GetOneJob(ctx context.Context, req *jobpb.GetAJobRequest) (*jobpb.JobOpeningResponse, error) {
	employerID := req.EmployerIDInt
	jobId := req.JobId

	res, err := js.jobUseCase.GetOneJob(employerID, jobId)
	if err != nil {
		return nil, err
	}
	salary := strconv.Itoa(res.Salary)
	jobOpening := &jobpb.JobOpeningResponse{
		Id:                  uint64(res.ID),
		Title:               res.Title,
		Description:         res.Description,
		Requirements:        res.Requirements,
		PostedOn:            timestamppb.New(res.PostedOn),
		Location:            res.Location,
		EmploymentType:      res.EmploymentType,
		Salary:              salary,
		SkillsRequired:      res.SkillsRequired,
		ExperienceLevel:     res.ExperienceLevel,
		EducationLevel:      res.EducationLevel,
		ApplicationDeadline: timestamppb.New(res.ApplicationDeadline),
		EmployerId:          employerID,
	}

	return jobOpening, nil
}

func (js *RecruiterJobServer) DeleteAJob(ctx context.Context, req *jobpb.DeleteAJobRequest) (*emptypb.Empty, error) {
	employerID := req.EmployerIDInt
	jobID := req.JobId

	err := js.jobUseCase.DeleteAJob(employerID, jobID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete job: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (js *RecruiterJobServer) UpdateAJob(ctx context.Context, req *jobpb.UpdateAJobRequest) (*jobpb.UpdateAJobResponse, error) {
	employerID := req.EmployerIDInt
	jobID := req.JobId

	jobDetails := models.JobOpening{
		Title:               req.Title,
		Description:         req.Description,
		Requirements:        req.Requirements,
		Location:            req.Location,
		EmploymentType:      req.EmploymentType,
		Salary:              req.Salary,
		SkillsRequired:      req.SkillsRequired,
		ExperienceLevel:     req.ExperienceLevel,
		EducationLevel:      req.EducationLevel,
		ApplicationDeadline: req.ApplicationDeadline.AsTime(),
	}

	fmt.Println("service", jobDetails)

	res, err := js.jobUseCase.UpdateAJob(employerID, jobID, jobDetails)
	if err != nil {
		return nil, err
	}
	salary := strconv.Itoa(res.Salary)
	updateResponse := &jobpb.UpdateAJobResponse{
		Id:                  uint64(res.ID),
		Title:               res.Title,
		Description:         res.Description,
		Requirements:        res.Requirements,
		PostedOn:            timestamppb.New(res.PostedOn),
		Location:            res.Location,
		EmploymentType:      res.EmploymentType,
		Salary:              salary,
		SkillsRequired:      res.SkillsRequired,
		ExperienceLevel:     res.ExperienceLevel,
		EducationLevel:      res.EducationLevel,
		ApplicationDeadline: timestamppb.New(res.ApplicationDeadline),
		EmployerId:          employerID,
	}

	return updateResponse, nil
}

func (js *RecruiterJobServer) GetJobAppliedCandidates(ctx context.Context, req *jobpb.GetAppliedJobsRequest) (*jobpb.GetAppliedJobsResponse, error) {
	employerID := req.UserId

	res, err := js.jobUseCase.GetJobAppliedCandidates(int(employerID))
	if err != nil {
		return nil, err
	}

	var jobs []*jobpb.AppliedJobs
	for _, job := range res.Jobs {
		jobs = append(jobs, &jobpb.AppliedJobs{
			JobId:       int64(job.JobID),
			Id:          int64(job.ID),
			UserId:      int64(job.JobseekerID),
			RecruiterId: int64(job.RecruiterID),
			Status:      job.Status,
			ResumeUrl:   job.ResumeUrl,
			CoverLetter: job.CoverLetter,
			Name:        job.JobseekerName,
			Email:       job.JoseekerEmail,
		})
	}
	response := &jobpb.GetAppliedJobsResponse{
		Jobs: jobs,
	}

	return response, nil
}

func (js *RecruiterJobServer) ScheduleInterview(ctx context.Context, req *jobpb.ScheduleInterviewRequest) (*jobpb.ScheduleInterviewResponse, error) {
	var data = models.ScheduleReq{
		ApplicationId: int(req.ApplicationId),
		RecruiterID:   uint(req.RecruiterId),
		DateAndTime:   req.DateAndTime.AsTime(),
		Mode:          req.Mode,
		Link:          req.Link,
	}

	res, err := js.jobUseCase.ScheduleInterview(data)
	if err != nil {
		return nil, err
	}

	return &jobpb.ScheduleInterviewResponse{
		Id:          int64(res.ID),
		JobId:       int64(res.JobID),
		JobseekerId: int64(res.JobseekerID),
		RecruiterId: req.RecruiterId,
		DateAndTime: req.DateAndTime,
		Mode:        res.Mode,
		Link:        res.Link,
		Status:      res.Status,
	}, nil
}
