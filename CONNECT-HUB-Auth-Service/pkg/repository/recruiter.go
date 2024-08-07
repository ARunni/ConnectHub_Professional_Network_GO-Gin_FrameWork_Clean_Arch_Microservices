package repository

import (
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/repository/interface"
	"github.com/ARunni/ConnetHub_auth/pkg/utils/models"
	req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type recruiterRepository struct {
	DB      *gorm.DB
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewRecruiterRepository(DB *gorm.DB) interfaces.RecruiterRepository {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &recruiterRepository{
		DB:      DB,
		Logger:  logger,
		LogFile: logFile,
	}
}
func (rr *recruiterRepository) RecruiterSignup(data req.RecruiterSignUp) (req.RecruiterDetailsResponse, error) {
	var recruiter req.RecruiterDetailsResponse
	querry := `insert into users 
	(company_name,industry,company_size,website,headquarters_address,about_company,contact_email,contact_phone_number,password)
	 values(?,?,?,?,?,?,?,?,?) RETURNING *`
	result := rr.DB.Raw(querry, data.Company_name,
		data.Industry, data.Company_size, data.Website,
		data.Headquarters_address, data.About_company,
		data.Contact_email, data.Contact_phone_number,
		data.Password).Scan(&recruiter)
	if result.Error != nil {
		return req.RecruiterDetailsResponse{}, result.Error
	}
	return recruiter, nil

}

func (rr *recruiterRepository) CheckRecruiterExistsByEmail(email string) (bool, error) {
	var count int
	querry := `select count(*) from users where contact_email = ? and role = ?`
	result := rr.DB.Raw(querry, email, "Recruiter").Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

// func (rr *recruiterRepository) CheckRecruiterExistsByEmail(email string) (bool, error) {
// 	var count int
// 	querry := `select count(*) from recruiters where contact_email = ?`
// 	result := rr.DB.Raw(querry, email).Scan(&count)
// 	if result.Error != nil {
// 		return false, result.Error
// 	}
// 	return count > 0, nil
// }

func (rr *recruiterRepository) RecruiterLogin(data req.RecruiterLogin) (req.RecruiterDetailsResponse, error) {
	var recruiter req.RecruiterDetailsResponse
	querry := ` select * from users where contact_email = ?`
	result := rr.DB.Raw(querry, data.Email).Scan(&recruiter)
	if result.Error != nil {
		return req.RecruiterDetailsResponse{}, result.Error
	}
	return recruiter, nil
}

// func (rr *recruiterRepository) RecruiterLogin(data req.RecruiterLogin) (req.RecruiterDetailsResponse, error) {
// 	var recruiter req.RecruiterDetailsResponse
// 	querry := ` select * from recruiters where contact_email = ?`
// 	result := rr.DB.Raw(querry, data.Email).Scan(&recruiter)
// 	if result.Error != nil {
// 		return req.RecruiterDetailsResponse{}, result.Error
// 	}
// 	return recruiter, nil
// }

func (rr *recruiterRepository) CheckRecruiterBlockByEmail(email string) (bool, error) {
	var ok bool
	err := rr.DB.Raw("select is_blocked from users where contact_email = ?", email).Scan(&ok).Error
	if err != nil {
		return false, err
	}
	return ok, nil
}

// func (rr *recruiterRepository) CheckRecruiterBlockByEmail(email string) (bool, error) {
// 	var ok bool
// 	err := rr.DB.Raw("select is_blocked from recruiters where contact_email = ?", email).Scan(&ok).Error
// 	if err != nil {
// 		return false, err
// 	}
// 	return ok, nil
// }

func (rr *recruiterRepository) RecruiterGetProfile(id int) (req.RecruiterProfile, error) {
	var recruiter req.RecruiterProfile

	querry := `select * from users where id = ?`

	result := rr.DB.Raw(querry, id).Scan(&recruiter)
	if result.Error != nil {
		return req.RecruiterProfile{}, result.Error
	}
	return recruiter, nil
}

// func (rr *recruiterRepository) RecruiterGetProfile(id int) (req.RecruiterProfile, error) {
// 	var recruiter req.RecruiterProfile

// 	querry := `select * from recruiters where id = ?`

// 	result := rr.DB.Raw(querry, id).Scan(&recruiter)
// 	if result.Error != nil {
// 		return req.RecruiterProfile{}, result.Error
// 	}
// 	return recruiter, nil
// }

func (rr *recruiterRepository) RecruiterEditProfile(profile req.RecruiterProfile) (req.RecruiterProfile, error) {
	p := profile

	querry := `update users set company_name =?,
	industry =?,company_size =?,website=?,headquarters_address=?,
	about_company=? ,contact_email =?,contact_phone_number = ?  where id = ?`

	result := rr.DB.Raw(querry, p.Company_name, p.Industry,
		p.Company_size, p.Website, p.Headquarters_address,
		p.About_company, p.Contact_email, p.Contact_phone_number, p.ID)
	if result.Error != nil {
		return req.RecruiterProfile{}, result.Error
	}
	return profile, nil
}

// func (rr *recruiterRepository) RecruiterEditProfile(profile req.RecruiterProfile) (req.RecruiterProfile, error) {
// 	p := profile

// 	querry := `update recruiters set company_name =?,
// 	industry =?,company_size =?,website=?,headquarters_address=?,
// 	about_company=? ,contact_email =?,contact_phone_number = ?  where id = ?`

// 	result := rr.DB.Raw(querry, p.Company_name, p.Industry,
// 		p.Company_size, p.Website, p.Headquarters_address,
// 		p.About_company, p.Contact_email, p.Contact_phone_number, p.ID)
// 	if result.Error != nil {
// 		return req.RecruiterProfile{}, result.Error
// 	}
// 	return profile, nil
// }

func (rr *recruiterRepository) IsRecruiterBlocked(id int) (bool, error) {
	var ok bool
	qurry := `select is_blocked from users where id = ?`
	err := rr.DB.Raw(qurry, id).Scan(&ok).Error
	if err != nil {
		return false, err
	}
	return ok, nil
}

// func (rr *recruiterRepository) IsRecruiterBlocked(id int) (bool, error) {
// 	var ok bool
// 	qurry := `select is_blocked from recruiters where id = ?`
// 	err := rr.DB.Raw(qurry, id).Scan(&ok).Error
// 	if err != nil {
// 		return false, err
// 	}
// 	return ok, nil
// }

// policies
func (rr *recruiterRepository) GetAllPolicies() (req.GetAllPolicyRes, error) {
	var pData []models.Policy
	qurry := `select id,title,content,created_at,updated_at from policies`
	err := rr.DB.Raw(qurry).Scan(&pData).Error
	if err != nil {
		return req.GetAllPolicyRes{}, err
	}
	return req.GetAllPolicyRes{Policies: pData}, nil
}

func (rr *recruiterRepository) GetOnePolicy(policy_id int) (req.CreatePolicyRes, error) {
	var pData models.Policy
	qurry := `select id,title,content,created_at,updated_at from policies where id = ?`
	err := rr.DB.Raw(qurry, policy_id).Scan(&pData).Error
	if err != nil {
		return req.CreatePolicyRes{}, err
	}
	return req.CreatePolicyRes{Policies: pData}, nil
}

func (rr *recruiterRepository) IsPolicyExist(policy_id int) (bool, error) {
	var data int
	qurry := `select count(*) from policies where id = ?`
	err := rr.DB.Raw(qurry, policy_id).Scan(&data).Error
	if err != nil {
		return false, err
	}
	return data > 0, nil
}
func (rr *recruiterRepository) GetDetailsById(userId int) (string, string, error) {
	var data models.UserData
	query := `SELECT email, first_name FROM users WHERE id = ? and role = ?`

	err := rr.DB.Raw(query, userId, "Jobseeker").Scan(&data).Error
	if err != nil {
		return "", "", err
	}

	return data.Email, data.FirstName, nil
}

// func (rr *recruiterRepository) GetDetailsById(userId int) (string, string, error) {
// 	var data models.UserData
// 	query := `SELECT email, first_name FROM job_seekers WHERE id = ?`

// 	err := rr.DB.Raw(query, userId).Scan(&data).Error
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return data.Email, data.FirstName, nil
// }

func (rr *recruiterRepository) GetDetailsByIdRecuiter(userId int) (string, string, error) {
	var data models.UserData
	query := `SELECT contact_email as email, company_name as first_name FROM users WHERE id = ? and role = ?`

	err := rr.DB.Raw(query, userId, "Recruiter").Scan(&data).Error
	if err != nil {
		return "", "", err
	}

	return data.Email, data.FirstName, nil
}

// func (rr *recruiterRepository) GetDetailsByIdRecuiter(userId int) (string, string, error) {
// 	var data models.UserData
// 	query := `SELECT contact_email as email, company_name as first_name FROM recruiters WHERE id = ?`

// 	err := rr.DB.Raw(query, userId).Scan(&data).Error
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return data.Email, data.FirstName, nil
// }
