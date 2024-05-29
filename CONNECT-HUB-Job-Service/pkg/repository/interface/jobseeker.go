package interfaces

import "ConnetHub_job/pkg/utils/models"

type JobseekerJobRepository interface {
	IsJobExist(jobID int32) (bool, error)
	JobSeekerGetAllJobs(keyword string) ([]models.JobOpeningData, error)
	JobSeekerGetJobByID(id int) (models.JobOpeningData, error)
	JobSeekerApplyJob(jobId, userId, recruiterId int) (bool, error)
	IsAppliedAlready(jobId, userId int) (bool, error)
	GetRecruiterByJobId(jobId int) (int, error)
	GetAppliedJobs(userId int) ([]models.ApplyJob, error)
}
