package interfaces

import "connectHub_gateway/pkg/utils/models"

type RecruiterJobClient interface {
	PostJob(data models.JobOpening) (models.JobOpeningData, error)
	GetAllJobs(recruiterID int32) ([]models.AllJob, error)
	GetOneJob(recruiterID, jobId int32) (models.JobOpeningData, error)
	DeleteAJob(recruiterID, jobId int32) error
}
