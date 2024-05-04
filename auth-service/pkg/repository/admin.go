package repository

import (
	interfaces "ConnetHub_auth/pkg/repository/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}
func (ar *adminRepository) CheckAdminExistsByEmail(email string) (bool, error) {
	qurry := `select count(*) from admins where email = ?`
	var count int
	err := ar.DB.Raw(qurry, email).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ar *adminRepository) AdminLogin(admin req.AdminLogin) (req.AdminDetailsResponse, error) {
	var adminDetails req.AdminDetailsResponse
	qurry := `select * from admins where email = ?`
	err := ar.DB.Raw(qurry, admin.Email).Scan(&adminDetails).Error
	if err != nil {
		return req.AdminDetailsResponse{}, err
	}
	return adminDetails, nil
}
