package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type JobSeekerUseCase interface {
	JobSeekerSignup(jobseekerdata req.JobSeekerSignUp) (req.TokenJobSeeker, error)
	JobSeekerLogin(jobseekerDetails req.JobSeekerLogin) (req.TokenJobSeeker, error)
	JobSeekerGetProfile(id int) (req.JobSeekerProfile, error)
}
