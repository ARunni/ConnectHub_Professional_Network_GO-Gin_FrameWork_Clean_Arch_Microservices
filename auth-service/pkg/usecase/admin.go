package usecase

import (
	"ConnetHub_auth/pkg/helper"
	interfaces "ConnetHub_auth/pkg/repository/interface"
	usecase "ConnetHub_auth/pkg/usecase/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) usecase.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails req.AdminLogin) (req.TokenAdmin, error) {

	adminCompareDetails, err := ad.adminRepository.AdminLogin(adminDetails)
	if err != nil {
		return req.TokenAdmin{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return req.TokenAdmin{}, err
	}
	var adminDetailsResponse req.AdminLoginRes
	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return req.TokenAdmin{}, err
	}
	access, err := helper.GenerateTokenAdmin(adminCompareDetails)
	if err != nil {
		return req.TokenAdmin{}, err
	}
	return req.TokenAdmin{
		Admin: adminCompareDetails, Token: access,
	}, nil
}
