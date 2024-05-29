package interfaces

import "ConnetHub_job/pkg/utils/models"

type JobseekerJobUsecase interface {
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
	JobSeekerGetJobByID(id int) (models.JobOpeningData, error)
	JobSeekerApplyJob(jobId, userId int) (bool, error)
	GetAppliedJobs(user_id int) (models.AppliedJobs, error)
}
