package interfaces

import "github.com/ARunni/connectHub_gateway/pkg/utils/models"

type RecruiterAuthClient interface {
	RecruiterSignup(recruiterData models.RecruiterSignUp) (models.TokenRecruiter, error)
	RecruiterLogin(recruiterData models.RecruiterLogin) (models.TokenRecruiter, error)
	RecruiterGetProfile(id int) (models.RecruiterProfile, error)
	RecruiterEditProfile(data models.RecruiterProfile) (models.RecruiterProfile, error)

	GetAllPolicies() (models.GetAllPolicyRes, error)
	GetOnePolicy(policy_id int) (models.CreatePolicyRes, error)
}
