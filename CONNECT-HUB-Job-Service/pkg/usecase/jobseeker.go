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
