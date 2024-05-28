package repository

import (
	interfaces "ConnetHub_auth/pkg/repository/interface"
	"ConnetHub_auth/pkg/utils/models"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"time"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}
func (ar *adminRepository) CheckAdminExistsByEmail(email string) (bool, error) {
	qurry := `select count(*) from admins where email = ?`
	var count int
	err := ar.DB.Raw(qurry, email).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ar *adminRepository) AdminLogin(admin req.AdminLogin) (req.AdminDetailsResponse, error) {
	var adminDetails req.AdminDetailsResponse
	qurry := `select * from admins where email = ?`
	err := ar.DB.Raw(qurry, admin.Email).Scan(&adminDetails).Error
	if err != nil {
		return req.AdminDetailsResponse{}, err
	}
	return adminDetails, nil
}

func (ar *adminRepository) GetRecruiters(page int) ([]req.RecruiterDetailsAtAdmin, error) {

	var recruiters []req.RecruiterDetailsAtAdmin

	// pagination purpose -
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * 2

	qurry := `select id,company_name,contact_email as contact_mail,contact_phone_number as phone,is_blocked as blocked from recruiters limit ? offset ?`
	err := ar.DB.Raw(qurry, 5, offset).Scan(&recruiters).Error
	if err != nil {
		return nil, err
	}
	return recruiters, nil
}

func (ar *adminRepository) GetJobseekers(page int) ([]req.JobseekerDetailsAtAdmin, error) {

	var jobseekers []req.JobseekerDetailsAtAdmin

	// pagination purpose -
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * 2

	qurry := `select id,first_name as name,email,phone_number as phone,is_blocked as blocked  from job_seekers limit ? offset ?`
	err := ar.DB.Raw(qurry, 5, offset).Scan(&jobseekers).Error
	if err != nil {
		return nil, err
	}
	return jobseekers, nil
}

func (ar *adminRepository) BlockRecruiter(id int) error {

	qurry := `update recruiters set is_blocked = true where id = ?`
	err := ar.DB.Exec(qurry, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *adminRepository) BlockJobseeker(id int) error {

	qurry := `update job_seekers set is_blocked = true where id = ?`
	err := ar.DB.Exec(qurry, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *adminRepository) UnBlockJobseeker(id int) error {

	qurry := `update job_seekers set is_blocked = false where id = ?`
	err := ar.DB.Exec(qurry, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *adminRepository) UnBlockRecruiter(id int) error {

	qurry := `update recruiters set is_blocked = false where id = ?`
	err := ar.DB.Exec(qurry, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *adminRepository) CheckJobseekerById(id int) (bool, error) {
	var count int
	qurry := `select count(*) from job_seekers where id = ?`
	err := ar.DB.Raw(qurry, id).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ar *adminRepository) CheckRecruiterById(id int) (bool, error) {
	var count int
	qurry := `select count(*) from recruiters where id = ?`
	err := ar.DB.Raw(qurry, id).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ar *adminRepository) IsJobseekerBlocked(id int) (bool, error) {
	var ok bool
	qurry := `select is_blocked from job_seekers where id = ?`
	err := ar.DB.Raw(qurry, id).Scan(&ok).Error
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (ar *adminRepository) IsRecruiterBlocked(id int) (bool, error) {
	var ok bool
	qurry := `select is_blocked from recruiters where id = ?`
	err := ar.DB.Raw(qurry, id).Scan(&ok).Error
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (ar *adminRepository) GetJobseekerDetails(id int) (req.JobseekerDetailsAtAdmin, error) {
	var data req.JobseekerDetailsAtAdmin
	qurry := `select id,first_name as name,email,phone_number as phone,is_blocked as blocked  from job_seekers where id = ?`
	err := ar.DB.Raw(qurry, id).Scan(&data).Error
	if err != nil {
		return req.JobseekerDetailsAtAdmin{}, err
	}
	return data, nil
}

func (ar *adminRepository) GetRecruiterDetails(id int) (req.RecruiterDetailsAtAdmin, error) {
	var data req.RecruiterDetailsAtAdmin
	qurry := `select id,company_name,contact_email as contact_mail,contact_phone_number as phone,is_blocked as blocked from recruiters where id = ?`
	err := ar.DB.Raw(qurry, id).Scan(&data).Error
	if err != nil {
		return req.RecruiterDetailsAtAdmin{}, err
	}
	return data, nil
}

// policies
func (ar *adminRepository) CreatePolicy(data req.CreatePolicyReq) (models.Policy, error) {

	var pData models.Policy

	qurry := `insert into policies 
	(title,content,created_at) 
	values ($1,$2,$3) returning id,title,content,created_at,updated_at`

	err := ar.DB.Raw(qurry, data.Title, data.Content, time.Now()).Scan(&pData).Error

	if err != nil {
		return models.Policy{}, err
	}

	return pData, nil
}

func (ar *adminRepository) UpdatePolicy(data req.UpdatePolicyReq) (models.Policy, error) {
	var pData models.Policy
	qurry := `update policies set title = $1,content = $2, updated_at =$3 where id = $4 returning id,title,content,created_at,updated_at`
	err := ar.DB.Raw(qurry, data.Title, data.Content, time.Now(), data.Id).Scan(&pData).Error
	if err != nil {
		return models.Policy{}, err
	}
	return pData, nil
}

func (ar *adminRepository) DeletePolicy(policy_id int) (bool, error) {

	qurry := `delete from policies where id = ?`
	err := ar.DB.Exec(qurry, policy_id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ar *adminRepository) GetAllPolicies() (req.GetAllPolicyRes, error) {

	var pData []models.Policy
	qurry := `select id,title,content,created_at,updated_at from policies`
	err := ar.DB.Raw(qurry).Scan(&pData).Error
	if err != nil {
		return req.GetAllPolicyRes{}, err
	}
	return req.GetAllPolicyRes{Policies: pData}, nil

}

func (ar *adminRepository) GetOnePolicy(policy_id int) (req.CreatePolicyRes, error) {
	var pData models.Policy
	qurry := `select id,title,content,created_at,updated_at from policies where id = ?`
	err := ar.DB.Raw(qurry, policy_id).Scan(&pData).Error
	if err != nil {
		return req.CreatePolicyRes{}, err
	}
	return req.CreatePolicyRes{Policies: pData}, nil
}

func (ar *adminRepository) IsPolicyExist(policy_id int) (bool, error) {
	var data int
	qurry := `select count(*) from policies where id = ?`
	err := ar.DB.Raw(qurry, policy_id).Scan(&data).Error
	if err != nil {
		return false, err
	}
	return data > 0, nil
}
