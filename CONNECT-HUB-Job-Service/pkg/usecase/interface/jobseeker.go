package interfaces

import "ConnetHub_job/pkg/utils/models"

type JobseekerJobUsecase interface {
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
	JobSeekerGetJobByID(id int) (models.JobOpeningData, error)

	GetAppliedJobs(user_id int) (models.AppliedJobs, error)
	JobSeekerApplyJob(data models.ApplyJobReq) (models.ApplyJob, error)
}
