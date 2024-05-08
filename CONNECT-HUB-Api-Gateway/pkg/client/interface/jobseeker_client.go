package interfaces

import "connectHub_gateway/pkg/utils/models"

type JobSeekerClient interface {
	JobSeekerSignup(jobseekerData models.JobSeekerSignUp) (models.TokenJobSeeker, error)
	JobSeekerLogin(jobseekerData models.JobSeekerLogin) (models.TokenJobSeeker, error)
}
