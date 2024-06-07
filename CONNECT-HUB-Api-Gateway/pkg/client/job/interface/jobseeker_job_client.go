package interfaces

import "connectHub_gateway/pkg/utils/models"

type JobseekerJobClient interface {
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
	JobSeekerGetJobByID(id int) (models.JobOpeningData, error)
	JobSeekerApplyJob(data models.ApplyJobReq) (models.ApplyJob, error)
	GetAppliedJobs(user_id int) (models.AppliedJobsJ, error) 
}
