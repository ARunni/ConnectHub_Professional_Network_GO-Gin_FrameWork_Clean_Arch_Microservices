package interfaces

import "connectHub_gateway/pkg/utils/models"

type JobseekerJobClient interface {
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
}
