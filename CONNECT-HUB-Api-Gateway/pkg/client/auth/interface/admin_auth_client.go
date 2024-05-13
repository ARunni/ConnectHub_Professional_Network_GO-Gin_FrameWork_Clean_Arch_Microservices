package interfaces

import "connectHub_gateway/pkg/utils/models"

type AdminAuthClient interface {
	AdminLogin(admindata models.AdminLogin) (models.TokenAdmin, error)
	GetRecruiters(page int) ([]models.RecruiterDetailsAtAdmin, error)
	GetJobseekers(page int) ([]models.JobseekerDetailsAtAdmin, error)
	BlockRecruiter(id int) (models.BlockRes, error)
	BlockJobseeker(id int) (models.BlockRes, error)
	UnBlockRecruiter(id int) (models.BlockRes, error)
	UnBlockJobseeker(id int) (models.BlockRes, error)
	GetJobseekerDetails(id int) (models.JobseekerDetailsAtAdmin, error)
	GetRecruiterDetails(id int) (models.RecruiterDetailsAtAdmin, error)
}
