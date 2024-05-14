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

	fmt.Println(jobSeekerJobs)

	return jobSeekerJobs, nil

}
