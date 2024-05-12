package repository

import (
	interfaces "ConnetHub_auth/pkg/repository/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"

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

	qurry := `select * from recruiters limit ? offset ?`
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

	qurry := `select * from job_seekers limit ? offset ?`
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
	qurry := `select * from job_seekers where id = ?`
	err := ar.DB.Raw(qurry, id).Scan(&data).Error
	if err != nil {
		return req.JobseekerDetailsAtAdmin{}, err
	}
	return data, nil
}

func (ar *adminRepository) GetRecruiterDetails(id int) (req.RecruiterDetailsAtAdmin, error) {
	var data req.RecruiterDetailsAtAdmin
	qurry := `select * from recruiters where id = ?`
	err := ar.DB.Raw(qurry, id).Scan(&data).Error
	if err != nil {
		return req.RecruiterDetailsAtAdmin{}, err
	}
	return data, nil
}
