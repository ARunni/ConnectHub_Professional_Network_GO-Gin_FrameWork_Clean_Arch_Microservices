package usecase

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	logging "github.com/ARunni/ConnetHub_job/Logging"
	"github.com/ARunni/ConnetHub_job/pkg/client/auth/interfaces"
	"github.com/ARunni/ConnetHub_job/pkg/helper"
	repo "github.com/ARunni/ConnetHub_job/pkg/repository/interface"
	usecase "github.com/ARunni/ConnetHub_job/pkg/usecase/interface"
	"github.com/ARunni/ConnetHub_job/pkg/utils/models"

	msg "github.com/ARunni/Error_Message"
	"github.com/sirupsen/logrus"
)

type recruiterJobUseCase struct {
	jobRepository repo.RecruiterJobRepository
	Client        interfaces.JobAuthClient
	Logger        *logrus.Logger
	LogFile       *os.File
}

func NewRecruiterJobUseCase(repo repo.RecruiterJobRepository, client interfaces.JobAuthClient) usecase.RecruiterJobUsecase {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Job.log")
	return &recruiterJobUseCase{
		jobRepository: repo,
		Client:        client,
		Logger:        logger,
		LogFile:       logFile,
	}
}

func (ju *recruiterJobUseCase) PostJob(data models.JobOpening) (models.JobOpeningData, error) {

	if data.EmployerID <= 0 {
		return models.JobOpeningData{}, errors.New(msg.ErrDataNegative)
	}
	if data.Title == "" || data.Description == "" ||
		data.Requirements == "" || data.Location == "" ||
		data.EmploymentType == "" || data.SkillsRequired == "" ||
		data.ExperienceLevel == "" || data.EducationLevel == "" {

		return models.JobOpeningData{}, errors.New(msg.ErrFieldEmpty)
	}

	date := time.Now()
	salary, err := strconv.Atoi(data.Salary)
	if err != nil {

		return models.JobOpeningData{}, errors.New(msg.ErrDatatypeConversion)
	}

	jobData := models.JobOpeningData{
		Title:               data.Title,
		Description:         data.Description,
		Requirements:        data.Requirements,
		PostedOn:            date,
		EmployerID:          data.EmployerID,
		Location:            data.Location,
		EmploymentType:      data.EmploymentType,
		Salary:              salary,
		SkillsRequired:      data.SkillsRequired,
		ExperienceLevel:     data.ExperienceLevel,
		EducationLevel:      data.EducationLevel,
		ApplicationDeadline: data.ApplicationDeadline,
	}
	jobDataRes, err := ju.jobRepository.PostJob(jobData)
	if err != nil {
		return models.JobOpeningData{}, err
	}
	return jobDataRes, nil
}

func (ju *recruiterJobUseCase) GetAllJobs(employerID int32) ([]models.AllJob, error) {

	jobData, err := ju.jobRepository.GetAllJobs(employerID)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}

func (ju *recruiterJobUseCase) GetOneJob(recruiterID, jobId int32) (models.JobOpeningData, error) {

	isJobExist, err := ju.jobRepository.IsJobExist(jobId)
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return models.JobOpeningData{}, fmt.Errorf("job with ID %d does not exist", jobId)
	}

	jobData, err := ju.jobRepository.GetOneJob(recruiterID, jobId)
	if err != nil {
		return models.JobOpeningData{}, err
	}
	return jobData, nil
}

func (ju *recruiterJobUseCase) DeleteAJob(employerIDInt, jobID int32) error {
	if employerIDInt <= 0 {
		return fmt.Errorf("employer ID must be greater than zero")
	}

	if jobID <= 0 {
		return fmt.Errorf("jobID ID must be greater than zero")
	}

	isJobExist, err := ju.jobRepository.IsJobExist(jobID)
	if err != nil {
		return fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return fmt.Errorf("job with ID %d does not exist", jobID)
	}

	// If the job exists, proceed with deletion
	err = ju.jobRepository.DeleteAJob(employerIDInt, jobID)
	if err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}

	return nil
}
func (ju *recruiterJobUseCase) UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningData, error) {
	fmt.Println("idddddddddddddd", employerID)

	if employerID <= 0 {
		return models.JobOpeningData{}, errors.New("recruiter id is not valid")
	}
	if jobID <= 0 {
		return models.JobOpeningData{}, errors.New("job id is not valid")
	}
	salaryInt, err := strconv.Atoi(jobDetails.Salary)
	if err != nil {
		return models.JobOpeningData{}, err
	}
	if salaryInt <= 0 {
		return models.JobOpeningData{}, errors.New("salary is not valid")
	}
	if jobDetails.Title == "" {
		return models.JobOpeningData{}, errors.New("title is required")
	}
	if jobDetails.Description == "" {
		return models.JobOpeningData{}, errors.New("description is required")
	}

	if jobDetails.Requirements == "" {
		return models.JobOpeningData{}, errors.New("requirements is required")
	}
	fmt.Println("jobdetailssssssss", jobDetails)
	fmt.Println("jobdetailsssseducationssss", jobDetails.EducationLevel)
	if jobDetails.EducationLevel == "" {
		return models.JobOpeningData{}, errors.New("educationLevel is required")
	}
	if jobDetails.EmploymentType == "" {
		return models.JobOpeningData{}, errors.New("employmentType is required")
	}
	if jobDetails.Location == "" {
		return models.JobOpeningData{}, errors.New("location is required")
	}

	if jobDetails.SkillsRequired == "" {
		return models.JobOpeningData{}, errors.New("skillsrequired is required")
	}

	isJobExist, err := ju.jobRepository.IsJobExist(jobID)
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return models.JobOpeningData{}, fmt.Errorf("job with ID %d does not exist", jobID)
	}

	// updation

	updatedJob, err := ju.jobRepository.UpdateAJob(employerID, jobID, jobDetails)
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to update job: %v", err)
	}

	return updatedJob, nil
}

func (ju *recruiterJobUseCase) GetJobAppliedCandidates(recruiter_id int) (models.AppliedJobs, error) {
	if recruiter_id <= 0 {
		return models.AppliedJobs{}, errors.New("recruiter id is not valid")
	}

	jobData, err := ju.jobRepository.GetJobAppliedCandidates(recruiter_id)
	if err != nil {
		return models.AppliedJobs{}, fmt.Errorf("failed to Get applied job: %v", err)
	}
	var jobDatas []models.ApplyJobRes
	for _, datas := range jobData {

		fmt.Println("datas", datas)
		email, name, err := ju.Client.GetDetailsById(int(datas.JobseekerID))
		if err != nil {
			return models.AppliedJobs{}, err
		}
		datas.JobseekerName = name
		datas.JoseekerEmail = email
		jobDatas = append(jobDatas, datas)

	}

	return models.AppliedJobs{Jobs: jobDatas}, nil
}

func (ju *recruiterJobUseCase) ChangeApplicationStatusToRejected(appId, recruiterID int) (bool, error) {
	if appId <= 0 {
		return false, errors.New("application id not valid")
	}
	okA, err := ju.jobRepository.ISApplicationExist(appId, recruiterID)
	if err != nil {
		return false, fmt.Errorf("failed to check if application exists: %v", err)
	}
	if !okA {
		return false, fmt.Errorf("application with ID %d does not exist or not belongs to you", appId)
	}

	jobData, err := ju.jobRepository.ChangeApplicationStatusToRejected(appId)
	if err != nil {
		return false, fmt.Errorf("failed to Get applied job: %v", err)
	}
	return jobData, nil
}

func (ju *recruiterJobUseCase) ScheduleInterview(data models.ScheduleReq) (models.Interview, error) {
	if data.ApplicationId <= 0 {
		return models.Interview{}, errors.New("application id not valid")
	}
	if data.Mode != "online" && data.Mode != "offline" {
		return models.Interview{}, errors.New("application mode should be online or offline")
	}
	if data.Link == "" {
		return models.Interview{}, errors.New("link is not valid")
	}
	okA, err := ju.jobRepository.ISApplicationExist(data.ApplicationId, int(data.RecruiterID))
	if err != nil {
		return models.Interview{}, fmt.Errorf("failed to check if application exists: %v", err)
	}
	if !okA {
		return models.Interview{}, fmt.Errorf("application with ID %d does not exist or not belongs to you", data.ApplicationId)
	}

	appData, err := ju.jobRepository.GetApplicationDetails(data.ApplicationId)
	if err != nil {
		return models.Interview{}, fmt.Errorf("failed to get application details: %v", err)
	}
	okI, err := ju.jobRepository.ISApplicationScheduled(data.ApplicationId)
	if err != nil {
		return models.Interview{}, fmt.Errorf("failed to check if application is scheduled: %v", err)
	}
	if okI {
		return models.Interview{}, fmt.Errorf("application with ID %d is already scheduled", data.ApplicationId)
	}

	var dataI = models.Interview{
		JobID:         appData.ID,
		JobseekerID:   appData.JobseekerID,
		RecruiterID:   data.RecruiterID,
		DateAndTime:   data.DateAndTime,
		Mode:          data.Mode,
		Link:          data.Link,
		ApplicationId: uint(data.ApplicationId),
	}

	jobData, err := ju.jobRepository.ScheduleInterview(dataI)
	if err != nil {
		return models.Interview{}, fmt.Errorf("failed to schedule interview: %v", err)
	}

	okR, err := ju.jobRepository.ChangeApplicationStatusToScheduled(data.ApplicationId)
	if err != nil {
		return models.Interview{}, fmt.Errorf("failed to Get applied job: %v", err)
	}
	if !okR {
		return models.Interview{}, fmt.Errorf("failed to schedule interview")
	}
	jobData.ApplicationId = uint(data.ApplicationId)

	// notification purpose
	_, recname, err := ju.Client.GetDetailsByIdRecuiter(int(data.RecruiterID))
	if err != nil {
		return models.Interview{}, fmt.Errorf("failed to initialize notification")
	}
	msg := fmt.Sprintf("%s Scheduled an Interview for Your Application Id %d", recname, data.ApplicationId)
	helper.SendNotification(models.Notification{
		UserID:     int(appData.JobseekerID),
		SenderID:   int(data.RecruiterID),
		PostID:     data.ApplicationId,
		SenderName: recname,
	}, []byte(msg))

	return jobData, nil
}

func (ju *recruiterJobUseCase) CancelScheduledInterview(appId, userId int) (bool, error) {
	if appId <= 0 {
		return false, errors.New("application id not valid")
	}

	okA, err := ju.jobRepository.ISApplicationExist(appId, userId)
	if err != nil {
		return false, fmt.Errorf("failed to check if application exists: %v", err)
	}
	if !okA {
		return false, fmt.Errorf("application with ID %d does not exist or not belongs to you", appId)
	}

	okI, err := ju.jobRepository.ISApplicationScheduled(appId)
	if err != nil {
		return false, fmt.Errorf("failed to check if application is scheduled: %v", err)
	}
	if !okI {
		return false, fmt.Errorf("application with ID %d is not scheduled", appId)
	}
	okS, err := ju.jobRepository.CancelScheduledApplication(appId)
	if err != nil {
		return false, fmt.Errorf("failed to Cancel scheduled interview: %v", err)
	}
	if !okS {
		return false, fmt.Errorf("failed to Cancel scheduled interview : %d", appId)
	}

	okR, err := ju.jobRepository.ChangeApplicationStatusToCancel(appId)
	if err != nil {
		return false, fmt.Errorf("failed to change status: %v", err)
	}
	if !okR {
		return false, fmt.Errorf("failed to cancel scheduled interview")
	}

	// notification purpose

	appData, err := ju.jobRepository.GetApplicationDetails(appId)
	if err != nil {
		return false, fmt.Errorf("failed to get application details: %v", err)
	}
	_, recname, err := ju.Client.GetDetailsByIdRecuiter(int(userId))
	if err != nil {
		return false, fmt.Errorf("failed to initialize notification")
	}
	msg := fmt.Sprintf("%s is canceled Scheduled Interview for Your Application Id %d", recname, appId)
	helper.SendNotification(models.Notification{
		UserID:     int(appData.JobseekerID),
		SenderID:   int(userId),
		PostID:     appId,
		SenderName: recname,
	}, []byte(msg))

	return true, nil
}
