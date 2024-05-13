package interfaces

import "ConnetHub_job/pkg/utils/models"

type JobRepository interface {
	PostJob(data models.JobOpeningData) (models.JobOpeningData, error)
}
