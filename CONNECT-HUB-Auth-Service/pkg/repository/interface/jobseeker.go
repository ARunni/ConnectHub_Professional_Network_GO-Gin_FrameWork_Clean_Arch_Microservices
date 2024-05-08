package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type JobseekerRepository interface {
	JobSeekerSignup(data req.JobSeekerSignUp) (req.JobSeekerDetailsResponse, error)
	JobseekerLogin(data req.JobSeekerLogin) (req.JobSeekerDetailsResponse, error)
	CheckJobseekerExistsByEmail(email string) (bool, error)
}
