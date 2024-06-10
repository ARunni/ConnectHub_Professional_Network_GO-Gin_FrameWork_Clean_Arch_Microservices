package interfaces

import req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"

type JobseekerRepository interface {
	JobSeekerSignup(data req.JobSeekerSignUp) (req.JobSeekerDetailsResponse, error)
	JobseekerLogin(data req.JobSeekerLogin) (req.JobSeekerDetailsResponse, error)
	CheckJobseekerExistsByEmail(email string) (bool, error)
	CheckJobseekerBlockByEmail(email string) (bool, error)
	JobSeekerGetProfile(id int) (req.JobSeekerProfile, error)
	JobSeekerEditProfile(data req.JobSeekerProfile) (req.JobSeekerProfile, error)
	IsJobseekerBlocked(id int) (bool, error)

	GetAllPolicies()(req.GetAllPolicyRes,error)
	GetOnePolicy(policy_id int) (req.CreatePolicyRes,error)
	IsPolicyExist(policy_id int)(bool,error)
}
