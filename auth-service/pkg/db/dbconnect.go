package db

import (
	"ConnetHub_auth/pkg/config"
	"ConnetHub_auth/pkg/utils/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.AutoMigrate(&models.Admin{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Recruiter{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.JobSeeker{}); err != nil {
		return nil, err
	}

	return db, nil
}
