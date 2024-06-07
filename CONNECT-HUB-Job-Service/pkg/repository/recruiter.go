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

func (jr *recruiterJobRepository) GetJobAppliedCandidates(recruiter_id int) ([]models.ApplyJobRes, error) {

	var jobs []models.ApplyJobRes
	if err := jr.DB.Raw("select * from apply_jobs where recruiter_id = ?", recruiter_id).Scan(&jobs).Error; err != nil {
		fmt.Println(err)
		return []models.ApplyJobRes{}, fmt.Errorf("failed to query jobs: %v", err)
	}

	return jobs, nil
}

// Scheluling interwiew
func (jr *recruiterJobRepository) ScheduleInterview(data models.Interview) (models.Interview, error) {

	var jobs models.Interview
	querry := `insert into interviews (job_id,jobseeker_id,recruiter_id,date_and_time,mode,link,application_id) 
	values ($1,$2,$3,$4,$5,$6,$7) returning id,job_id,jobseeker_id,recruiter_id,date_and_time,mode,link,status,application_id`
	if err := jr.DB.Raw(querry, data.JobID, data.JobseekerID, data.RecruiterID, data.DateAndTime, data.Mode, data.Link, data.ApplicationId).Scan(&jobs).Error; err != nil {
		return models.Interview{}, fmt.Errorf("failed to query schedule interview: %v", err)
	}

	return jobs, nil
}

func (jr *recruiterJobRepository) ISApplicationExist(appId, recruiterId int) (bool, error) {

	var data int
	if err := jr.DB.Raw("select count(*) from apply_jobs where recruiter_id = ? and id = ?", recruiterId, appId).Scan(&data).Error; err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("failed to query jobs: %v", err)
	}

	return data > 0, nil
}

func (jr *recruiterJobRepository) ISApplicationScheduled(appId int) (bool, error) {

	var data int
	if err := jr.DB.Raw("select count(*) from interviews where application_id = ?", appId).Scan(&data).Error; err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("failed to query jobs: %v", err)
	}

	return data > 0, nil
}

func (jr *recruiterJobRepository) GetApplicationDetails(appId int) (models.ApplyJob, error) {

	var jobs models.ApplyJob
	if err := jr.DB.Raw("select * from apply_jobs where id = ?", appId).Scan(&jobs).Error; err != nil {
		fmt.Println(err)
		return models.ApplyJob{}, fmt.Errorf("failed to query jobs: %v", err)
	}

	return jobs, nil
}

func (jr *recruiterJobRepository) ChangeApplicationStatusToScheduled(appId int) (bool, error) {

	if err := jr.DB.Exec("update apply_jobs set status = ? where id = ?", "scheduled", appId).Error; err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("failed to query jobs: %v", err)
	}

	return true, nil
}

func (jr *recruiterJobRepository) ChangeApplicationStatusToRejected(appId int) (bool, error) {

	if err := jr.DB.Raw("update apply_jobs set status = ? where id = ?", "rejected", appId).Error; err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("failed to query jobs: %v", err)
	}

	return true, nil
}
