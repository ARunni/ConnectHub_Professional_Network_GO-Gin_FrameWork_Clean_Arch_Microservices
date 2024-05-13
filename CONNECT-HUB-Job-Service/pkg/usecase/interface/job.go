package interfaces

import "ConnetHub_job/pkg/utils/models"

type JobUsecase interface {
	PostJob(data models.JobOpening) (models.JobOpeningData, error)
}
