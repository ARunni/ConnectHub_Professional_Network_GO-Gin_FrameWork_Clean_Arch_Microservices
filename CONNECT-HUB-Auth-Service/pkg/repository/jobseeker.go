package repository

import (
	interfaces "ConnetHub_auth/pkg/repository/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"

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

func (jr *jobseekerRepository) JobseekerSignup(data req.JobSeekerSignUp) (req.JobSeekerDetailsResponse, error) {

	var res req.JobSeekerDetailsResponse

	querry := `insert into jobseekers 
	(email,password,first_name,last_name,phone_number,date_of_birth,gender,created_at)
	 values(?,?,?,?,?,?,?,NOW()) RETURNING id`
	result := jr.DB.Raw(querry, data.Email, data.FirstName, data.LastName, data.PhoneNumber, data.DateOfBirth, data.Gender).Scan(&res)

	if result.Error != nil {
		return req.JobSeekerDetailsResponse{}, result.Error
	}
	return res, nil
}

func (jr *jobseekerRepository) JobseekerLogin(data req.JobSeekerLogin) (req.JobSeekerDetailsResponse, error) {
	var res req.JobSeekerDetailsResponse
	querry := ` select * from jobseekers where email = ?`
	result := jr.DB.Raw(querry, data.Email).Scan(&res)
	if result.Error != nil {
		return req.JobSeekerDetailsResponse{}, result.Error
	}
	return res, nil
}

func (jr *jobseekerRepository) CheckJobseekerExistsByEmail(email string) (bool, error) {
	var count int
	querry := `select count(*) from jobseekers where email = ?`
	result := jr.DB.Raw(querry, email).Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
