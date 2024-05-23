package usecase

import (
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

func (ju *jobseekerJobUseCase) JobSeekerApplyJob(jobId, userId int) (bool, error) {

	jobOk, err := ju.jobRepository.JobSeekerApplyJob(jobId,userId)
	if err != nil {
		return false, fmt.Errorf("failed to apply job: %v", err)
	}

	return jobOk, nil
}
