package usecase

import (
	repo "ConnetHub_job/pkg/repository/interface"
	interfaces "ConnetHub_job/pkg/usecase/interface"
	"ConnetHub_job/pkg/utils/models"
	"errors"
	"strconv"
	"time"

	msg "github.com/ARunni/Error_Message"
)

type jobUseCase struct {
	jobRepository repo.JobRepository
}

func NewJobUseCase(repo repo.JobRepository) interfaces.JobUsecase {
	return &jobUseCase{
		jobRepository: repo,
	}
}

func (ju *jobUseCase) PostJob(data models.JobOpening) (models.JobOpeningData, error) {

	if data.EmployerID <= 0 {
		return models.JobOpeningData{}, errors.New(msg.ErrDataNegative)
	}
	if data.Title == "" || data.Description == "" ||
		data.Requirements == "" || data.Location == "" ||
		data.EmploymentType == "" || data.SkillsRequired == "" ||
		data.ExperienceLevel == "" || data.EducationLevel == "" {

		return models.JobOpeningData{}, errors.New(msg.ErrFieldEmpty)
	}

	date := time.Now()
	salary, err := strconv.Atoi(data.Salary)
	if err != nil {

		return models.JobOpeningData{}, errors.New(msg.ErrDatatypeConversion)
	}

	jobData := models.JobOpeningData{
		Title:               data.Title,
		Description:         data.Description,
		Requirements:        data.Requirements,
		PostedOn:            date,
		EmployerID:          data.EmployerID,
		Location:            data.Location,
		EmploymentType:      data.EmploymentType,
		Salary:              salary,
		SkillsRequired:      data.SkillsRequired,
		ExperienceLevel:     data.ExperienceLevel,
		EducationLevel:      data.EducationLevel,
		ApplicationDeadline: data.ApplicationDeadline,
	}
	jobDataRes, err := ju.jobRepository.PostJob(jobData)
	if err != nil {
		return models.JobOpeningData{}, err
	}
	return jobDataRes, nil
}
