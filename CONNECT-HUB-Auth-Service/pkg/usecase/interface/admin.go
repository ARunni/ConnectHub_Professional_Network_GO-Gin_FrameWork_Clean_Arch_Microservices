package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type AdminUseCase interface {
	AdminLogin(adminDetails req.AdminLogin) (req.TokenAdmin, error)
	GetRecruiters(page int) ([]req.RecruiterDetailsAtAdmin, error)
	GetJobseekers(page int) ([]req.JobseekerDetailsAtAdmin, error)
	BlockRecruiter(id int) (req.BlockRes, error)
	BlockJobseeker(id int) (req.BlockRes, error)
	UnBlockRecruiter(id int) (req.BlockRes, error)
	UnBlockJobseeker(id int) (req.BlockRes, error)
}
