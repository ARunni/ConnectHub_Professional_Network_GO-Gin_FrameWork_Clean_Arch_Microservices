package db

import (
	"ConnetHub_auth/pkg/config"
	"ConnetHub_auth/pkg/utils/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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
	CheckAndCreateAdmin(db)
	return db, nil
}

func CheckAndCreateAdmin(db *gorm.DB) {
	var count int64
	db.Model(&models.Admin{}).Count(&count)
	if count == 0 {
		password := "admin@123"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		admin := models.Admin{
			ID:        1,
			Firstname: "ConnectHub",
			Lastname:  "Admin",
			Email:     "admin@connecthub.com",
			Password:  string(hashedPassword),
		}
		db.Create(&admin)
		fmt.Println("Admin Created")
	}
}
