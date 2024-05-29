package interfaces

import "ConnetHub_job/pkg/utils/models"

type RecruiterJobRepository interface {
	PostJob(data models.JobOpeningData) (models.JobOpeningData, error)
	GetAllJobs(recruiterID int32) ([]models.AllJob, error)
	GetOneJob(recruiterID, jobId int32) (models.JobOpeningData, error)
	DeleteAJob(recruiterID, jobId int32) error
	IsJobExist(jobID int32) (bool, error)
	UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningData, error)
	GetJobAppliedCandidates(recruiter_id int) ([]models.ApplyJob, error)
}
