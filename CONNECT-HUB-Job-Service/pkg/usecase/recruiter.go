package usecase

import (
	repo "ConnetHub_job/pkg/repository/interface"
	interfaces "ConnetHub_job/pkg/usecase/interface"
	"ConnetHub_job/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"
	"time"

	msg "github.com/ARunni/Error_Message"
)

type recruiterJobUseCase struct {
	jobRepository repo.RecruiterJobRepository
}

func NewRecruiterJobUseCase(repo repo.RecruiterJobRepository) interfaces.RecruiterJobUsecase {
	return &recruiterJobUseCase{
		jobRepository: repo,
	}
}

func (ju *recruiterJobUseCase) PostJob(data models.JobOpening) (models.JobOpeningData, error) {
	fmt.Println("recruiter id ", data.EmployerID)
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

	isJobExist, err := ju.jobRepository.IsJobExist(jobID)
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return models.JobOpeningData{}, fmt.Errorf("job with ID %d does not exist", jobID)
	}

	updatedJob, err := ju.jobRepository.UpdateAJob(employerID, jobID, jobDetails)
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to update job: %v", err)
	}

	return updatedJob, nil
}

func (ju *recruiterJobUseCase) GetJobAppliedCandidates(recruiter_id int) (models.AppliedJobs, error) {

	jobData, err := ju.jobRepository.GetJobAppliedCandidates(recruiter_id)
	if err != nil {
		return models.AppliedJobs{}, fmt.Errorf("failed to Get applied job: %v", err)
	}
	return models.AppliedJobs{Jobs: jobData}, nil
}
