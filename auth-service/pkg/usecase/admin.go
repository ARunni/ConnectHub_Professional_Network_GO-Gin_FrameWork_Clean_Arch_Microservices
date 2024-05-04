package usecase

import (
	"ConnetHub_auth/pkg/helper"
	interfaces "ConnetHub_auth/pkg/repository/interface"
	usecase "ConnetHub_auth/pkg/usecase/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"errors"

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

func (au *adminUseCase) LoginHandler(adminDetails req.AdminLogin) (req.TokenAdmin, error) {
	ok, err := au.adminRepository.CheckAdminExistsByEmail(adminDetails.Email)
	if err != nil {
		return req.TokenAdmin{}, err
	}
	if !ok {
		return req.TokenAdmin{}, errors.New("no admin found")
	}

	adminCompareDetails, err := au.adminRepository.AdminLogin(adminDetails)
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
