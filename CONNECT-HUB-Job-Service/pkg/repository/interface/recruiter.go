package interfaces

import "ConnetHub_job/pkg/utils/models"

type RecruiterJobRepository interface {
	PostJob(data models.JobOpeningData) (models.JobOpeningData, error)
}
