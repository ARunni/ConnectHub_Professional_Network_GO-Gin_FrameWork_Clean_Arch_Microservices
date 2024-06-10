package interfaces

import "github.com/ARunni/ConnetHub_job/pkg/utils/models"

type RecruiterJobRepository interface {
	PostJob(data models.JobOpeningData) (models.JobOpeningData, error)
	GetAllJobs(recruiterID int32) ([]models.AllJob, error)
	GetOneJob(recruiterID, jobId int32) (models.JobOpeningData, error)
	DeleteAJob(recruiterID, jobId int32) error
	IsJobExist(jobID int32) (bool, error)
	UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningData, error)
	GetJobAppliedCandidates(recruiter_id int) ([]models.ApplyJobRes, error)

	ScheduleInterview(data models.Interview) (models.Interview, error)
	ISApplicationExist(appId, recruiterId int) (bool, error)
	GetApplicationDetails(appId int) (models.ApplyJob, error)
	ISApplicationScheduled(appId int) (bool, error)
	ChangeApplicationStatusToScheduled(appId int) (bool, error)
	ChangeApplicationStatusToRejected(appId int) (bool, error)
}
