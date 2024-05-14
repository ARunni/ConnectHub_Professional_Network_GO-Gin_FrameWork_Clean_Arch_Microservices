package interfaces

import "ConnetHub_job/pkg/utils/models"

type JobseekerJobUsecase interface {
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
}
