package interfaces

import "github.com/ARunni/ConnetHub_job/pkg/utils/models"

type JobseekerJobUsecase interface {
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
	JobSeekerGetJobByID(id int) (models.JobOpeningData, error)

	GetAppliedJobs(user_id int) (models.AppliedJobsJ, error)
	JobSeekerApplyJob(data models.ApplyJobReq) (models.ApplyJob, error)

	GetInterviewDetails(appId, userId int) (models.InterviewResp, error)
}
