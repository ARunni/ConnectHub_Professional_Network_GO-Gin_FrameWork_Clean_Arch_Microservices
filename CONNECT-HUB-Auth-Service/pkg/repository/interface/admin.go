package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type AdminRepository interface {
	AdminLogin(admin req.AdminLogin) (req.AdminDetailsResponse, error)
	CheckAdminExistsByEmail(email string) (bool, error)
	GetRecruiters(page int) ([]req.RecruiterDetailsAtAdmin, error)
	GetJobseekers(page int) ([]req.JobseekerDetailsAtAdmin, error)
	BlockRecruiter(id int) error
	BlockJobseeker(id int) error
	UnBlockRecruiter(id int) error
	UnBlockJobseeker(id int) error
	CheckJobseekerById(id int) (bool, error)
	CheckRecruiterById(id int) (bool, error)
	IsJobseekerBlocked(id int) (bool, error)
	IsRecruiterBlocked(id int) (bool, error)
	GetJobseekerDetails(id int) (req.JobseekerDetailsAtAdmin, error)
	GetRecruiterDetails(id int) (req.RecruiterDetailsAtAdmin, error)
}
