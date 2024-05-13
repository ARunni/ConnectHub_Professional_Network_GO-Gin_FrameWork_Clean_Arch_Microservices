package interfaces

import "connectHub_gateway/pkg/utils/models"

type JobSeekerAuthClient interface {
	JobSeekerSignup(jobseekerData models.JobSeekerSignUp) (models.TokenJobSeeker, error)
	JobSeekerLogin(jobseekerData models.JobSeekerLogin) (models.TokenJobSeeker, error)
	JobSeekerGetProfile(id int) (models.JobSeekerProfile, error)
	JobSeekerEditProfile(profile models.JobSeekerProfile) (models.JobSeekerProfile, error)
}
