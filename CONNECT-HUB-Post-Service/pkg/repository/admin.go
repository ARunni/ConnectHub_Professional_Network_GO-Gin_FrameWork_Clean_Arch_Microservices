package repository

import (
	"ConnetHub_post/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type adminPostRepository struct {
	DB *gorm.DB
}

func NewAdminPostRepository(DB *gorm.DB) interfaces.AdminPostRepository {
	return &adminPostRepository{
		DB: DB,
	}
}
