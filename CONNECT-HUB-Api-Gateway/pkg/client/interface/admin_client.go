package interfaces

import "connectHub_gateway/pkg/utils/models"

type AdminClient interface {
	AdminLogin(admindata models.AdminLogin) (models.TokenAdmin, error)
	GetRecruiters(page int) ([]models.RecruiterDetailsAtAdmin, error)
	GetJobseekers(page int) ([]models.JobseekerDetailsAtAdmin, error)
	BlockRecruiter(id int) (models.BlockRes, error)
	BlockJobseeker(id int) (models.BlockRes, error)
	UnBlockRecruiter(id int) (models.BlockRes, error)
	UnBlockJobseeker(id int) (models.BlockRes, error)
}
