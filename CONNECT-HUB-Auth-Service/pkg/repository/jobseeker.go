package repository

import (
	interfaces "ConnetHub_auth/pkg/repository/interface"
	"ConnetHub_auth/pkg/utils/models"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"time"

	"gorm.io/gorm"
)

type jobseekerRepository struct {
	DB *gorm.DB
}

func NewJobseekerRepository(DB *gorm.DB) interfaces.JobseekerRepository {
	return &jobseekerRepository{
		DB: DB,
	}
}

func (jr *jobseekerRepository) JobSeekerSignup(data req.JobSeekerSignUp) (req.JobSeekerDetailsResponse, error) {

	var res req.JobSeekerDetailsResponse
	time := time.Now()
	input := models.JobSeeker{
		Email:         data.Email,
		Password:      data.Password,
		First_name:    data.FirstName,
		Last_name:     data.LastName,
		Phone_number:  data.PhoneNumber,
		Date_of_birth: data.DateOfBirth,
		Gender:        data.Gender,
		Created_at:    time,
	}
	result := jr.DB.Create(&input).Scan(&res)

	if result.Error != nil {
		return req.JobSeekerDetailsResponse{}, result.Error
	}
	return res, nil
}

func (jr *jobseekerRepository) JobseekerLogin(data req.JobSeekerLogin) (req.JobSeekerDetailsResponse, error) {
	var res req.JobSeekerDetailsResponse
	querry := ` select * from job_seekers where email = ?`
	result := jr.DB.Raw(querry, data.Email).Scan(&res)
	if result.Error != nil {
		return req.JobSeekerDetailsResponse{}, result.Error
	}
	return res, nil
}

func (jr *jobseekerRepository) CheckJobseekerExistsByEmail(email string) (bool, error) {
	var count int
	querry := `select count(*) from job_seekers where email = ?`
	result := jr.DB.Raw(querry, email).Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
