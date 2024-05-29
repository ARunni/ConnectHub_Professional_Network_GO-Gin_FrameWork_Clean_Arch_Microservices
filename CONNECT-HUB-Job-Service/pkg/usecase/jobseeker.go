package usecase

import (
	"ConnetHub_job/pkg/config"
	"ConnetHub_job/pkg/helper"
	repo "ConnetHub_job/pkg/repository/interface"
	interfaces "ConnetHub_job/pkg/usecase/interface"
	"ConnetHub_job/pkg/utils/models"
	"fmt"
)

type jobseekerJobUseCase struct {
	jobRepository repo.JobseekerJobRepository
}

func NewJobseekerJobUseCase(repo repo.JobseekerJobRepository) interfaces.JobseekerJobUsecase {
	return &jobseekerJobUseCase{
		jobRepository: repo,
	}
}

func (ju *jobseekerJobUseCase) JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error) {

	jobs, err := ju.jobRepository.JobSeekerGetAllJobs(keyword)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs: %v", err)
	}

	var jobSeekerJobs []models.JobSeekerGetAllJobs
	for _, job := range jobs {

		jobSeekerJob := models.JobSeekerGetAllJobs{
			ID:    job.ID,
			Title: job.Title,
		}
		jobSeekerJobs = append(jobSeekerJobs, jobSeekerJob)
	}

	return jobSeekerJobs, nil
}

func (ju *jobseekerJobUseCase) JobSeekerGetJobByID(id int) (models.JobOpeningData, error) {

	ok, err := ju.jobRepository.IsJobExist(int32(id))
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to check if job exist: %v", err)
	}
	if !ok {
		return models.JobOpeningData{}, fmt.Errorf("job not found")
	}

	job, err := ju.jobRepository.JobSeekerGetJobByID(id)
	if err != nil {
		return models.JobOpeningData{}, fmt.Errorf("failed to get job: %v", err)
	}

	return job, nil
}

func (ju *jobseekerJobUseCase) JobSeekerApplyJob(data models.ApplyJobReq) (models.ApplyJob, error) {
	if data.JobID <= 0 {
		return models.ApplyJob{}, fmt.Errorf("invalid job id")
	}
	ok, err := ju.jobRepository.IsJobExist(int32(data.JobID))
	if err != nil {
		return models.ApplyJob{}, fmt.Errorf("failed to check if job exist: %v", err)
	}
	if !ok {
		return models.ApplyJob{}, fmt.Errorf("job not found")
	}
	applyOk, err := ju.jobRepository.IsAppliedAlready(int(data.JobID), int(data.JobseekerID))
	if err != nil {
		return models.ApplyJob{}, fmt.Errorf("failed to check if already applied: %v", err)
	}
	if applyOk {
		return models.ApplyJob{}, fmt.Errorf("already applied")
	}
	recruiterId, err := ju.jobRepository.GetRecruiterByJobId(int(data.JobID))
	if err != nil {
		return models.ApplyJob{}, fmt.Errorf("failed to get recruiter id: %v", err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		return models.ApplyJob{}, err
	}

	h := helper.NewHelper(cfg)
	// coverUrl, err := h.AddImageToAwsS3(data.CoverLetter)
	// if err != nil {
	// 	return models.ApplyJob{}, err
	// }

	resumeUrl, err := h.AddImageToAwsS3(data.Resume)
	if err != nil {
		return models.ApplyJob{}, err
	}

	var reqJob = models.ApplyJob{
		JobID:       data.JobID,
		JobseekerID: data.JobseekerID,
		RecruiterID: uint(recruiterId),
		CoverLetter: data.CoverLetter,
		ResumeUrl:   resumeUrl,
	}

	jobOk, err := ju.jobRepository.JobSeekerApplyJob(reqJob)
	if err != nil {
		return models.ApplyJob{}, fmt.Errorf("failed to apply job: %v", err)
	}

	return jobOk, nil
}

func (ju *jobseekerJobUseCase) GetAppliedJobs(user_id int) (models.AppliedJobs, error) {

	jobData, err := ju.jobRepository.GetAppliedJobs(user_id)
	if err != nil {
		return models.AppliedJobs{}, fmt.Errorf("failed to Get applied job: %v", err)
	}

	return models.AppliedJobs{Jobs: jobData}, nil
}
