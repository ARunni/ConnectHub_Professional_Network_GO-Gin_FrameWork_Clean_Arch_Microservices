package interfaces

import "connectHub_gateway/pkg/utils/models"

type JobClient interface {
	PostJob(data models.JobOpening) (models.JobOpeningData, error)
}
