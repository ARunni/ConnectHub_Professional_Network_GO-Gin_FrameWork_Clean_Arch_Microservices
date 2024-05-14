package interfaces

import "ConnetHub_job/pkg/utils/models"

type JobseekerJobRepository interface {
	// IsJobExist(jobID int32) (bool, error)
	JobSeekerGetAllJobs(keyword string) ([]models.JobOpeningData, error)
}
