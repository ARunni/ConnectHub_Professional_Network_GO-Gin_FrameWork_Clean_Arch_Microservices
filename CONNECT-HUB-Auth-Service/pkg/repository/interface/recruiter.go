package interfaces

import req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"

type RecruiterRepository interface {
	RecruiterSignup(data req.RecruiterSignUp) (req.RecruiterDetailsResponse, error)
	RecruiterLogin(data req.RecruiterLogin) (req.RecruiterDetailsResponse, error)
	CheckRecruiterExistsByEmail(email string) (bool, error)
	CheckRecruiterBlockByEmail(email string) (bool, error)
	IsRecruiterBlocked(id int) (bool, error)
	RecruiterGetProfile(id int) (req.RecruiterProfile, error)
	RecruiterEditProfile(profile req.RecruiterProfile) (req.RecruiterProfile, error)

	GetAllPolicies() (req.GetAllPolicyRes, error)
	GetOnePolicy(policy_id int) (req.CreatePolicyRes, error)
	IsPolicyExist(policy_id int) (bool, error)

	GetDetailsById(userId int) (string, string, error)
	GetDetailsByIdRecuiter(userId int) (string, string, error)
}
