package repository

import (
	"os"
	"time"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/repository/interface"
	"github.com/ARunni/ConnetHub_auth/pkg/utils/models"
	req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB      *gorm.DB
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &adminRepository{
		DB:      DB,
		Logger:  logger,
		LogFile: logFile,
	}
}
func (ar *adminRepository) CheckAdminExistsByEmail(email string) (bool, error) {
	ar.Logger.Info("CheckAdminExistsByEmail at adminRepository started ")
	qurry := `select count(*) from admins where email = ?`
	var count int
	err := ar.DB.Raw(qurry, email).Scan(&count).Error
	if err != nil {
		ar.Logger.Error("Error in CheckAdminExistsByEmail at adminRepository : ", err)
		return false, err
	}
	ar.Logger.Info("CheckAdminExistsByEmail at adminRepository success ")
	return count > 0, nil
}

func (ar *adminRepository) AdminLogin(admin req.AdminLogin) (req.AdminDetailsResponse, error) {
	ar.Logger.Info("AdminLogin at adminRepository started ")
	var adminDetails req.AdminDetailsResponse
	qurry := `select * from admins where email = ?`
	err := ar.DB.Raw(qurry, admin.Email).Scan(&adminDetails).Error
	if err != nil {
		ar.Logger.Error("Error in AdminLogin at adminRepository : ", err)
		return req.AdminDetailsResponse{}, err
	}
	ar.Logger.Info("AdminLogin at adminRepository success ")
	return adminDetails, nil
}

func (ar *adminRepository) GetRecruiters(page int) ([]req.RecruiterDetailsAtAdmin, error) {
	ar.Logger.Info("GetRecruiters at adminRepository started ")

	var recruiters []req.RecruiterDetailsAtAdmin

	// pagination purpose -
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * 2

	qurry := `select id,company_name,contact_email as contact_mail,contact_phone_number as phone,is_blocked as blocked from users where role = ? limit ? offset ?`
	err := ar.DB.Raw(qurry,"Recruiter", 5, offset).Scan(&recruiters).Error
	if err != nil {
		ar.Logger.Error("Error in GetRecruiters at adminRepository : ", err)
		return nil, err
	}
	ar.Logger.Info("GetRecruiters at adminRepository success ")
	return recruiters, nil
}

func (ar *adminRepository) GetJobseekers(page int) ([]req.JobseekerDetailsAtAdmin, error) {

	ar.Logger.Info("GetJobseekers at adminRepository started ")

	var jobseekers []req.JobseekerDetailsAtAdmin

	// pagination purpose -
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * 2

	qurry := `select id,first_name as name,email,phone_number as phone,is_blocked as blocked  from users where role = ? limit ? offset ?`
	err := ar.DB.Raw(qurry,"Jobseeker", 5, offset).Scan(&jobseekers).Error
	if err != nil {
		ar.Logger.Error("Error in GetJobseekers at adminRepository : ", err)
		return nil, err
	}
	ar.Logger.Info("GetJobseekers at adminRepository success ")
	return jobseekers, nil
}
// func (ar *adminRepository) GetJobseekers(page int) ([]req.JobseekerDetailsAtAdmin, error) {

// 	ar.Logger.Info("GetJobseekers at adminRepository started ")

// 	var jobseekers []req.JobseekerDetailsAtAdmin

// 	// pagination purpose -
// 	if page == 0 {
// 		page = 1
// 	}

// 	offset := (page - 1) * 2

// 	qurry := `select id,first_name as name,email,phone_number as phone,is_blocked as blocked  from job_seekers limit ? offset ?`
// 	err := ar.DB.Raw(qurry, 5, offset).Scan(&jobseekers).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in GetJobseekers at adminRepository : ", err)
// 		return nil, err
// 	}
// 	ar.Logger.Info("GetJobseekers at adminRepository success ")
// 	return jobseekers, nil
// }

func (ar *adminRepository) BlockRecruiter(id int) error {
	ar.Logger.Info("BlockRecruiter at adminRepository started ")

	qurry := `update users set is_blocked = true where id = ? and role = ?`
	err := ar.DB.Exec(qurry, id,"Recruiter").Error
	if err != nil {
		ar.Logger.Error("Error in BlockRecruiter at adminRepository : ", err)
		return err
	}
	ar.Logger.Info("BlockRecruiter at adminRepository success ")
	return nil
}
// func (ar *adminRepository) BlockRecruiter(id int) error {
// 	ar.Logger.Info("BlockRecruiter at adminRepository started ")

// 	qurry := `update recruiters set is_blocked = true where id = ?`
// 	err := ar.DB.Exec(qurry, id).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in BlockRecruiter at adminRepository : ", err)
// 		return err
// 	}
// 	ar.Logger.Info("BlockRecruiter at adminRepository success ")
// 	return nil
// }

func (ar *adminRepository) BlockJobseeker(id int) error {
	ar.Logger.Info("BlockJobseeker at adminRepository started ")

	qurry := `update users set is_blocked = true where id = ? and role = ?`
	err := ar.DB.Exec(qurry, id,"Jobseeker").Error
	if err != nil {
		ar.Logger.Error("Error in BlockJobseeker at adminRepository : ", err)
		return err
	}
	ar.Logger.Info("BlockJobseeker at adminRepository success ")
	return nil
}
// func (ar *adminRepository) BlockJobseeker(id int) error {
// 	ar.Logger.Info("BlockJobseeker at adminRepository started ")

// 	qurry := `update job_seekers set is_blocked = true where id = ?`
// 	err := ar.DB.Exec(qurry, id).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in BlockJobseeker at adminRepository : ", err)
// 		return err
// 	}
// 	ar.Logger.Info("BlockJobseeker at adminRepository success ")
// 	return nil
// }

func (ar *adminRepository) UnBlockJobseeker(id int) error {

	ar.Logger.Info("UnBlockJobseeker at adminRepository started ")

	qurry := `update users set is_blocked = false where id = ? and role = ?`
	err := ar.DB.Exec(qurry, id,"Jobseeker").Error
	if err != nil {
		ar.Logger.Error("Error in UnBlockJobseeker at adminRepository : ", err)
		return err
	}
	ar.Logger.Info("UnBlockJobseeker at adminRepository success ")
	return nil
}
// func (ar *adminRepository) UnBlockJobseeker(id int) error {

// 	ar.Logger.Info("UnBlockJobseeker at adminRepository started ")

// 	qurry := `update job_seekers set is_blocked = false where id = ?`
// 	err := ar.DB.Exec(qurry, id).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in UnBlockJobseeker at adminRepository : ", err)
// 		return err
// 	}
// 	ar.Logger.Info("UnBlockJobseeker at adminRepository success ")
// 	return nil
// }

func (ar *adminRepository) UnBlockRecruiter(id int) error {

	ar.Logger.Info("UnBlockJobseeker at adminRepository started ")

	qurry := `update users set is_blocked = false where id = ? and role = ?`
	err := ar.DB.Exec(qurry, id,"Recruiter").Error
	if err != nil {
		ar.Logger.Error("Error in UnBlockRecruiter at adminRepository : ", err)
		return err
	}
	ar.Logger.Info("UnBlockRecruiter at adminRepository success ")
	return nil
}
// func (ar *adminRepository) UnBlockRecruiter(id int) error {

// 	ar.Logger.Info("UnBlockJobseeker at adminRepository started ")

// 	qurry := `update recruiters set is_blocked = false where id = ?`
// 	err := ar.DB.Exec(qurry, id).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in UnBlockRecruiter at adminRepository : ", err)
// 		return err
// 	}
// 	ar.Logger.Info("UnBlockRecruiter at adminRepository success ")
// 	return nil
// }

func (ar *adminRepository) CheckJobseekerById(id int) (bool, error) {
	ar.Logger.Info("CheckJobseekerById at adminRepository started ")
	var count int
	qurry := `select count(*) from users where id = ? and role = ?`
	err := ar.DB.Raw(qurry, id,"Jobseeker").Scan(&count).Error
	if err != nil {
		ar.Logger.Error("Error in CheckJobseekerById at adminRepository : ", err)
		return false, err
	}
	ar.Logger.Info("CheckJobseekerById at adminRepository success ")
	return count > 0, nil
}
// func (ar *adminRepository) CheckJobseekerById(id int) (bool, error) {
// 	ar.Logger.Info("CheckJobseekerById at adminRepository started ")
// 	var count int
// 	qurry := `select count(*) from job_seekers where id = ?`
// 	err := ar.DB.Raw(qurry, id).Scan(&count).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in CheckJobseekerById at adminRepository : ", err)
// 		return false, err
// 	}
// 	ar.Logger.Info("CheckJobseekerById at adminRepository success ")
// 	return count > 0, nil
// }

func (ar *adminRepository) CheckRecruiterById(id int) (bool, error) {
	ar.Logger.Info("CheckRecruiterById at adminRepository started ")
	var count int
	qurry := `select count(*) from users where id = ? and role = ?`
	err := ar.DB.Raw(qurry, id,"Recruiter").Scan(&count).Error
	if err != nil {
		ar.Logger.Error("Error in CheckRecruiterById at adminRepository : ", err)
		return false, err
	}
	ar.Logger.Info("CheckRecruiterById at adminRepository success ")
	return count > 0, nil
}
// func (ar *adminRepository) CheckRecruiterById(id int) (bool, error) {
// 	ar.Logger.Info("CheckRecruiterById at adminRepository started ")
// 	var count int
// 	qurry := `select count(*) from recruiters where id = ?`
// 	err := ar.DB.Raw(qurry, id).Scan(&count).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in CheckRecruiterById at adminRepository : ", err)
// 		return false, err
// 	}
// 	ar.Logger.Info("CheckRecruiterById at adminRepository success ")
// 	return count > 0, nil
// }

func (ar *adminRepository) IsJobseekerBlocked(id int) (bool, error) {
	ar.Logger.Info("IsJobseekerBlocked at adminRepository started ")
	var ok bool
	qurry := `select is_blocked from users where id = ? and role = ?`
	err := ar.DB.Raw(qurry, id,"Jobseeker").Scan(&ok).Error
	if err != nil {
		ar.Logger.Error("Error in IsJobseekerBlocked at adminRepository : ", err)
		return false, err
	}
	ar.Logger.Info("IsJobseekerBlocked at adminRepository success ")
	return ok, nil
}
// func (ar *adminRepository) IsJobseekerBlocked(id int) (bool, error) {
// 	ar.Logger.Info("IsJobseekerBlocked at adminRepository started ")
// 	var ok bool
// 	qurry := `select is_blocked from job_seekers where id = ?`
// 	err := ar.DB.Raw(qurry, id).Scan(&ok).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in IsJobseekerBlocked at adminRepository : ", err)
// 		return false, err
// 	}
// 	ar.Logger.Info("IsJobseekerBlocked at adminRepository success ")
// 	return ok, nil
// }

func (ar *adminRepository) IsRecruiterBlocked(id int) (bool, error) {
	ar.Logger.Info("IsRecruiterBlocked at adminRepository started ")
	var ok bool
	qurry := `select is_blocked from users where id = ? and role = ?`
	err := ar.DB.Raw(qurry, id,"Recruiter").Scan(&ok).Error
	if err != nil {
		ar.Logger.Error("Error in IsRecruiterBlocked at adminRepository : ", err)
		return false, err
	}
	ar.Logger.Info("IsRecruiterBlocked at adminRepository success ")
	return ok, nil
}
// func (ar *adminRepository) IsRecruiterBlocked(id int) (bool, error) {
// 	ar.Logger.Info("IsRecruiterBlocked at adminRepository started ")
// 	var ok bool
// 	qurry := `select is_blocked from recruiters where id = ?`
// 	err := ar.DB.Raw(qurry, id).Scan(&ok).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in IsRecruiterBlocked at adminRepository : ", err)
// 		return false, err
// 	}
// 	ar.Logger.Info("IsRecruiterBlocked at adminRepository success ")
// 	return ok, nil
// }

func (ar *adminRepository) GetJobseekerDetails(id int) (req.JobseekerDetailsAtAdmin, error) {
	ar.Logger.Info("GetJobseekerDetails at adminRepository started ")
	var data req.JobseekerDetailsAtAdmin
	qurry := `select id,first_name as name,email,phone_number as phone,is_blocked as blocked  from users where id = ? and role = ?`
	err := ar.DB.Raw(qurry, id,"Jobseeker").Scan(&data).Error
	if err != nil {
		ar.Logger.Error("Error in GetJobseekerDetails at adminRepository : ", err)
		return req.JobseekerDetailsAtAdmin{}, err
	}
	ar.Logger.Info("GetJobseekerDetails at adminRepository success ")
	return data, nil
}
// func (ar *adminRepository) GetJobseekerDetails(id int) (req.JobseekerDetailsAtAdmin, error) {
// 	ar.Logger.Info("GetJobseekerDetails at adminRepository started ")
// 	var data req.JobseekerDetailsAtAdmin
// 	qurry := `select id,first_name as name,email,phone_number as phone,is_blocked as blocked  from job_seekers where id = ?`
// 	err := ar.DB.Raw(qurry, id).Scan(&data).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in GetJobseekerDetails at adminRepository : ", err)
// 		return req.JobseekerDetailsAtAdmin{}, err
// 	}
// 	ar.Logger.Info("GetJobseekerDetails at adminRepository success ")
// 	return data, nil
// }

func (ar *adminRepository) GetRecruiterDetails(id int) (req.RecruiterDetailsAtAdmin, error) {
	ar.Logger.Info("GetRecruiterDetails at adminRepository started ")
	var data req.RecruiterDetailsAtAdmin
	qurry := `select id,company_name,contact_email as contact_mail,contact_phone_number as phone,is_blocked as blocked from users where id = ? and role = ?`
	err := ar.DB.Raw(qurry, id,"Recruiter").Scan(&data).Error
	if err != nil {
		ar.Logger.Error("Error in GetRecruiterDetails at adminRepository : ", err)
		return req.RecruiterDetailsAtAdmin{}, err
	}
	ar.Logger.Info("GetRecruiterDetails at adminRepository success ")
	return data, nil
}
// func (ar *adminRepository) GetRecruiterDetails(id int) (req.RecruiterDetailsAtAdmin, error) {
// 	ar.Logger.Info("GetRecruiterDetails at adminRepository started ")
// 	var data req.RecruiterDetailsAtAdmin
// 	qurry := `select id,company_name,contact_email as contact_mail,contact_phone_number as phone,is_blocked as blocked from recruiters where id = ?`
// 	err := ar.DB.Raw(qurry, id).Scan(&data).Error
// 	if err != nil {
// 		ar.Logger.Error("Error in GetRecruiterDetails at adminRepository : ", err)
// 		return req.RecruiterDetailsAtAdmin{}, err
// 	}
// 	ar.Logger.Info("GetRecruiterDetails at adminRepository success ")
// 	return data, nil
// }

// policies
func (ar *adminRepository) CreatePolicy(data req.CreatePolicyReq) (models.Policy, error) {

	ar.Logger.Info("CreatePolicy at adminRepository started ")

	var pData models.Policy

	qurry := `insert into policies 
	(title,content,created_at) 
	values ($1,$2,$3) returning id,title,content,created_at,updated_at`

	err := ar.DB.Raw(qurry, data.Title, data.Content, time.Now()).Scan(&pData).Error

	if err != nil {
		ar.Logger.Error("Error in CreatePolicy at adminRepository : ", err)
		return models.Policy{}, err
	}
	ar.Logger.Info("CreatePolicy at adminRepository success ")

	return pData, nil
}

func (ar *adminRepository) UpdatePolicy(data req.UpdatePolicyReq) (models.Policy, error) {
	ar.Logger.Info("UpdatePolicy at adminRepository started ")
	var pData models.Policy
	qurry := `update policies set title = $1,content = $2, updated_at =$3 where id = $4 returning id,title,content,created_at,updated_at`
	err := ar.DB.Raw(qurry, data.Title, data.Content, time.Now(), data.Id).Scan(&pData).Error
	if err != nil {
		ar.Logger.Error("Error in UpdatePolicy at adminRepository : ", err)
		return models.Policy{}, err
	}
	ar.Logger.Info("UpdatePolicy at adminRepository success ")
	return pData, nil
}

func (ar *adminRepository) DeletePolicy(policy_id int) (bool, error) {
	ar.Logger.Info("DeletePolicy at adminRepository started ")

	qurry := `delete from policies where id = ?`
	err := ar.DB.Exec(qurry, policy_id).Error
	if err != nil {
		ar.Logger.Error("Error in DeletePolicy at adminRepository : ", err)
		return false, err
	}
	ar.Logger.Info("DeletePolicy at adminRepository success ")
	return true, nil
}

func (ar *adminRepository) GetAllPolicies() (req.GetAllPolicyRes, error) {
	ar.Logger.Info("GetAllPolicies at adminRepository started ")

	var pData []models.Policy
	qurry := `select id,title,content,created_at,updated_at from policies`
	err := ar.DB.Raw(qurry).Scan(&pData).Error
	if err != nil {
		ar.Logger.Error("Error in GetAllPolicies at adminRepository : ", err)
		return req.GetAllPolicyRes{}, err
	}
	ar.Logger.Info("GetAllPolicies at adminRepository success ")
	return req.GetAllPolicyRes{Policies: pData}, nil

}

func (ar *adminRepository) GetOnePolicy(policy_id int) (req.CreatePolicyRes, error) {
	ar.Logger.Info("GetOnePolicy at adminRepository started ")
	var pData models.Policy
	qurry := `select id,title,content,created_at,updated_at from policies where id = ?`
	err := ar.DB.Raw(qurry, policy_id).Scan(&pData).Error
	if err != nil {
		ar.Logger.Error("Error in GetOnePolicy at adminRepository : ", err)
		return req.CreatePolicyRes{}, err
	}
	ar.Logger.Info("GetOnePolicy at adminRepository success ")
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
