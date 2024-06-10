package interfaces

import req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"

type AdminUseCase interface {
	AdminLogin(adminDetails req.AdminLogin) (req.TokenAdmin, error)
	GetRecruiters(page int) ([]req.RecruiterDetailsAtAdmin, error)
	GetJobseekers(page int) ([]req.JobseekerDetailsAtAdmin, error)
	BlockRecruiter(id int) (req.BlockRes, error)
	BlockJobseeker(id int) (req.BlockRes, error)
	UnBlockRecruiter(id int) (req.BlockRes, error)
	UnBlockJobseeker(id int) (req.BlockRes, error)
	GetJobseekerDetails(id int) (req.JobseekerDetailsAtAdmin, error)
	GetRecruiterDetails(id int) (req.RecruiterDetailsAtAdmin, error)

	CreatePolicy(data req.CreatePolicyReq) (req.CreatePolicyRes, error)
	UpdatePolicy(data req.UpdatePolicyReq) (req.CreatePolicyRes, error)
	DeletePolicy(policy_id int) (bool, error)
	GetAllPolicies() (req.GetAllPolicyRes, error)
	GetOnePolicy(policy_id int) (req.CreatePolicyRes,error)
}
