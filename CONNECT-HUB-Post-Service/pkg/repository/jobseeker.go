package repository

import (
	"ConnetHub_post/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type jobseekerPostRepository struct {
	DB *gorm.DB
}

func NewJobseekerPostRepository(DB *gorm.DB) interfaces.JobseekerPostRepository {
	return &jobseekerPostRepository{
		DB: DB,
	}
}
