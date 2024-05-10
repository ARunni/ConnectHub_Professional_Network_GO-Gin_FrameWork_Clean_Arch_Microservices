package usecase

import (
	"ConnetHub_auth/pkg/helper"
	interfaces "ConnetHub_auth/pkg/repository/interface"
	usecase "ConnetHub_auth/pkg/usecase/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"errors"
	"fmt"

	msg "github.com/ARunni/Error_Message"
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

func (au *adminUseCase) AdminLogin(adminDetails req.AdminLogin) (req.TokenAdmin, error) {
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

func (au *adminUseCase) GetRecruiters(page int) ([]req.RecruiterDetailsAtAdmin, error) {
	if page == 0 {
		return nil, errors.New(msg.ErrDataZero)
	}
	recruiters, err := au.adminRepository.GetRecruiters(page)
	if err != nil {
		return nil, err
	}
	return recruiters, nil

}
func (au *adminUseCase) GetJobseekers(page int) ([]req.JobseekerDetailsAtAdmin, error) {
	if page <= 0 {
		return nil, errors.New(msg.ErrDataZero)
	}
	jobseeker, err := au.adminRepository.GetJobseekers(page)
	if err != nil {
		return nil, err
	}
	return jobseeker, nil

}

func (au *adminUseCase) BlockRecruiter(id int) (req.BlockRes, error) {
	if id <= 0 {
		return req.BlockRes{}, errors.New(msg.ErrDataZero)
	}
	ok, err := au.adminRepository.CheckRecruiterById(id)
	if err != nil {
		return req.BlockRes{}, err
	}
	if !ok {
		return req.BlockRes{}, errors.New(msg.ErrIdExist)
	}

	ok, err = au.adminRepository.IsRecruiterBlocked(id)
	if err != nil {
		return req.BlockRes{}, err
	}
	if ok {
		return req.BlockRes{}, errors.New(msg.ErrBlockAlready)
	}

	err = au.adminRepository.BlockRecruiter(id)

	if err != nil {
		return req.BlockRes{}, err
	}
	return req.BlockRes{Status: "Success"}, nil

}

func (au *adminUseCase) BlockJobseeker(id int) (req.BlockRes, error) {
	if id <= 0 {
		return req.BlockRes{}, errors.New(msg.ErrDataZero)
	}
	ok, err := au.adminRepository.CheckJobseekerById(id)
	if err != nil {
		return req.BlockRes{}, err
	}

	if !ok {
		return req.BlockRes{}, errors.New(msg.ErrIdExist)
	}
	okk, err := au.adminRepository.IsJobseekerBlocked(id)
	if err != nil {

		return req.BlockRes{}, err
	}
	fmt.Println("jahgkj", okk)
	if okk {

		return req.BlockRes{}, errors.New(msg.ErrBlockAlready)
	}

	err = au.adminRepository.BlockJobseeker(id)
	if err != nil {
		return req.BlockRes{}, err
	}
	return req.BlockRes{Status: "Success"}, nil

}

func (au *adminUseCase) UnBlockRecruiter(id int) (req.BlockRes, error) {
	if id <= 0 {
		return req.BlockRes{}, errors.New(msg.ErrDataZero)
	}

	ok, err := au.adminRepository.CheckRecruiterById(id)
	if err != nil {
		return req.BlockRes{}, err
	}
	if !ok {
		return req.BlockRes{}, errors.New(msg.ErrIdExist)
	}
	okk, err := au.adminRepository.IsRecruiterBlocked(id)
	if err != nil {
		return req.BlockRes{}, err
	}
	if !okk {
		return req.BlockRes{}, errors.New(msg.ErrUnBlockAlready)
	}
	err = au.adminRepository.UnBlockRecruiter(id)
	if err != nil {
		return req.BlockRes{}, err
	}
	return req.BlockRes{Status: "Success"}, nil

}

func (au *adminUseCase) UnBlockJobseeker(id int) (req.BlockRes, error) {
	if id <= 0 {
		return req.BlockRes{}, errors.New(msg.ErrDataZero)
	}
	ok, err := au.adminRepository.CheckJobseekerById(id)
	if err != nil {
		return req.BlockRes{}, err
	}

	if !ok {
		return req.BlockRes{}, errors.New(msg.ErrIdExist)
	}
	okk, err := au.adminRepository.IsJobseekerBlocked(id)
	if err != nil {
		return req.BlockRes{}, err
	}

	if !okk {
		return req.BlockRes{}, errors.New(msg.ErrUnBlockAlready)
	}

	err = au.adminRepository.UnBlockJobseeker(id)
	if err != nil {
		return req.BlockRes{}, err
	}
	return req.BlockRes{Status: "Success"}, nil

}
