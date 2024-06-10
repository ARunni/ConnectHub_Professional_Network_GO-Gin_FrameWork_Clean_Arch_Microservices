package interfaces

import req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"

type JobSeekerUseCase interface {
	JobSeekerSignup(jobseekerdata req.JobSeekerSignUp) (req.TokenJobSeeker, error)
	JobSeekerLogin(jobseekerDetails req.JobSeekerLogin) (req.TokenJobSeeker, error)
	JobSeekerGetProfile(id int) (req.JobSeekerProfile, error)
	JobSeekerEditProfile(data req.JobSeekerProfile) (req.JobSeekerProfile, error)

	GetAllPolicies()(req.GetAllPolicyRes,error)
	GetOnePolicy(policy_id int) (req.CreatePolicyRes,error)
	
}
