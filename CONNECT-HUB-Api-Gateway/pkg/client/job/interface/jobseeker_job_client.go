package interfaces

import "connectHub_gateway/pkg/utils/models"

type JobseekerJobClient interface {
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
	JobSeekerGetJobByID(id int) (models.JobOpeningData, error)
	JobSeekerApplyJob(jobId, userId int) (bool, error)
	GetAppliedJobs(user_id int) (models.AppliedJobs, error) 
}
