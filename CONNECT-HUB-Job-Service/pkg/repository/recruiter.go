package repository

import (
	interfaces "ConnetHub_job/pkg/repository/interface"
	"ConnetHub_job/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type recruiterJobRepository struct {
	DB *gorm.DB
}

func NewRecruiterJobRepository(DB *gorm.DB) interfaces.RecruiterJobRepository {
	return &recruiterJobRepository{
		DB: DB,
	}
}

func (jr *recruiterJobRepository) PostJob(data models.JobOpeningData) (models.JobOpeningData, error) {

	if err := jr.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil

}

func (jr *recruiterJobRepository) GetAllJobs(recruiterID int32) ([]models.AllJob, error) {
	var jobs []models.AllJob

	if err := jr.DB.Model(&models.JobOpeningData{}).Select("id, title, application_deadline, employer_id").Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (jr *recruiterJobRepository) GetOneJob(recruiterID, jobId int32) (models.JobOpeningData, error) {
	var job models.JobOpeningData

	if err := jr.DB.Model(&models.JobOpeningData{}).Where("id = ? AND employer_id = ?", jobId, recruiterID).First(&job).Error; err != nil {
		return models.JobOpeningData{}, err
	}

	return job, nil
}

func (jr *recruiterJobRepository) IsJobExist(jobID int32) (bool, error) {
	var job models.JobOpeningData

	if err := jr.DB.Model(&models.JobOpeningData{}).Where("id = ?", jobID).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
func (jr *recruiterJobRepository) DeleteAJob(employerIDInt, jobID int32) error {

	if err := jr.DB.Delete(&models.JobOpeningData{}, jobID).Error; err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}

	return nil
}

func (jr *recruiterJobRepository) UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningData, error) {

	postedOn := time.Now()
	salary, err := strconv.Atoi(jobDetails.Salary)
	if err != nil {
		return models.JobOpeningData{}, err
	}
	updatedJob := models.JobOpeningData{
		ID:                  uint(jobID),
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		PostedOn:            postedOn,
		EmployerID:          int(employerID),
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: jobDetails.ApplicationDeadline,
	}
	fmt.Println("job id ", updatedJob.ID)
	fmt.Println("all struct ", updatedJob)

	if err := jr.DB.Model(&models.JobOpeningData{}).Where("id = ?", jobID).Updates(updatedJob).Error; err != nil {
		return models.JobOpeningData{}, err
	}
	fmt.Println("updated employment type", updatedJob.EmploymentType)

	return updatedJob, nil
}

func (jr *recruiterJobRepository) GetJobAppliedCandidates(recruiter_id int) ([]models.ApplyJob, error) {

	var jobs []models.ApplyJob
	if err := jr.DB.Raw("select * from apply_jobs where recruiter_id = ?", recruiter_id).Scan(&jobs).Error; err != nil {
		fmt.Println(err)
		return []models.ApplyJob{}, fmt.Errorf("failed to query jobs: %v", err)
	}

	return jobs, nil
}
