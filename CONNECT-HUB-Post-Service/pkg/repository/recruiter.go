package repository

import (
	"ConnetHub_post/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type recruiterPostRepository struct {
	DB *gorm.DB
}

func NewRecruiterPostRepository(DB *gorm.DB) interfaces.RecruiterPostRepository {
	return &recruiterPostRepository{
		DB: DB,
	}
}
