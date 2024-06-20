package interfaces

import "github.com/ARunni/ConnetHub_job/pkg/utils/models"

type RecruiterJobUsecase interface {
	PostJob(data models.JobOpening) (models.JobOpeningData, error)
	GetAllJobs(recruiterID int32) ([]models.AllJob, error)
	GetOneJob(recruiterID, jobId int32) (models.JobOpeningData, error)
	DeleteAJob(recruiterID, jobId int32) error
	UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningData, error)
	GetJobAppliedCandidates(recruiter_id int)(models.AppliedJobs, error)

	ScheduleInterview(data models.ScheduleReq) (models.Interview,error)
	CancelScheduledInterview(appId, userId int) (bool,error)
	ChangeApplicationStatusToRejected(appId,recruiterID int) (bool,error)

}
