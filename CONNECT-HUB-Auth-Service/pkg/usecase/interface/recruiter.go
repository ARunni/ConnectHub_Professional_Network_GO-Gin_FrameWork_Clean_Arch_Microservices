package interfaces

import req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"

type RecruiterUseCase interface {
	RecruiterSignup(recruiterdata req.RecruiterSignUp) (req.TokenRecruiter, error)
	RecruiterLogin(recruiterDetails req.RecruiterLogin) (req.TokenRecruiter, error)
	RecruiterGetProfile(id int) (req.RecruiterProfile, error)
	RecruiterEditProfile(profile req.RecruiterProfile) (req.RecruiterProfile, error)

	GetAllPolicies()(req.GetAllPolicyRes,error)
	GetOnePolicy(policy_id int) (req.CreatePolicyRes,error)

	GetDetailsById(userId int) (string, string, error)
	GetDetailsByIdRecuiter(userId int) (string, string, error)
}
