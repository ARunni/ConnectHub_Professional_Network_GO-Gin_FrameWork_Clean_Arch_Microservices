package repository

import (
	interfaces "ConnetHub_auth/pkg/repository/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"

	"gorm.io/gorm"
)

type recruiterRepository struct {
	DB *gorm.DB
}

func NewRecruiterRepository(DB *gorm.DB) interfaces.RecruiterRepository {
	return &recruiterRepository{
		DB: DB,
	}
}
func (ar *recruiterRepository) RecruiterSignup(data req.RecruiterSignUp) (req.RecruiterDetailsResponse, error) {
	var recruiter req.RecruiterDetailsResponse
	querry := `insert into recruiters 
	(company_name,industry,company_size,website,headquarters_address,about_company,contact_email,contact_phone_number,password)
	 values(?,?,?,?,?,?,?,?,?)`
	result := ar.DB.Raw(querry, data.Company_name,
		data.Industry, data.Company_size, data.Website,
		data.Headquarters_address, data.About_company,
		data.Contact_email, data.Contact_phone_number,
		data.Password).Scan(&recruiter)
	if result.Error != nil {
		return req.RecruiterDetailsResponse{}, result.Error
	}
	return recruiter, nil

}

func (ar *recruiterRepository) CheckRecruiterExistsByEmail(email string) (bool, error) {
	var count int
	querry := `select count(*) from recruiters where contact_email = ?`
	result := ar.DB.Raw(querry, email).Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (ar *recruiterRepository) RecruiterLogin(data req.RecruiterLogin) (req.RecruiterDetailsResponse, error) {
	var recruiter req.RecruiterDetailsResponse
	querry := ` select * from recruiters where contact_email = ?`
	result := ar.DB.Raw(querry, data.Email).Scan(&recruiter)
	if result.Error != nil {
		return req.RecruiterDetailsResponse{}, result.Error
	}
	return recruiter, nil
}
