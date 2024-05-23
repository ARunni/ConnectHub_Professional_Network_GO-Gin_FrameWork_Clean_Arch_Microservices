package repository

import (
	interfaces "ConnetHub_job/pkg/repository/interface"
	"ConnetHub_job/pkg/utils/models"
	"fmt"

	"gorm.io/gorm"
)

type jobseekerJobRepository struct {
	DB *gorm.DB
}

func NewjobseekerJobRepository(DB *gorm.DB) interfaces.JobseekerJobRepository {
	return &jobseekerJobRepository{
		DB: DB,
	}
}

func (jr *jobseekerJobRepository) JobSeekerGetAllJobs(keyword string) ([]models.JobOpeningData, error) {
	var jobSeekerJobs []models.JobOpeningData

	if err := jr.DB.Where("title ILIKE ?", "%"+keyword+"%").Find(&jobSeekerJobs).Error; err != nil {
		return nil, fmt.Errorf("failed to query jobs: %v", err)
	}

	return jobSeekerJobs, nil

}

func (jr *jobseekerJobRepository) JobSeekerGetJobByID(id int) (models.JobOpeningData, error) {
	var jobSeekerJob models.JobOpeningData

	if err := jr.DB.Raw("select * from job_opening_data where id = ? ", id).Scan(&jobSeekerJob).Error; err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to query jobs: %v", err)
	}

	return jobSeekerJob, nil

}

func (jr *jobseekerJobRepository) IsJobExist(jobID int32) (bool, error) {
	var count int

	if err := jr.DB.Raw("select count(*) from job_opening_data where id = ? ", jobID).Scan(&count).Error; err != nil {
		return false, fmt.Errorf("failed to query jobs: %v", err)
	}

	return count > 0, nil

}

func (jr *jobseekerJobRepository) JobSeekerApplyJob(jobId, userId int) (bool, error) {

	if err := jr.DB.Exec("insert into apply_jobs (job_id,jobseeker_id) values (?,?)", jobId, userId).Error; err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("failed to query jobs: %v", err)
	}

	return true, nil

}
