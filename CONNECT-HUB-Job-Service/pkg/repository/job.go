package repository

import (
	interfaces "ConnetHub_job/pkg/repository/interface"
	"ConnetHub_job/pkg/utils/models"

	"gorm.io/gorm"
)

type jobRepository struct {
	DB *gorm.DB
}

func NewJobRepository(DB *gorm.DB) interfaces.JobRepository {
	return &jobRepository{
		DB: DB,
	}
}

func (jr *jobRepository) PostJob(data models.JobOpeningData) (models.JobOpeningData, error) {

	if err := jr.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil

}
