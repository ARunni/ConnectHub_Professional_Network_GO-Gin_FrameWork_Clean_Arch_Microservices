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

func (au *adminUseCase) GetJobseekerDetails(id int) (req.JobseekerDetailsAtAdmin, error) {

	if id <= 0 {
		return req.JobseekerDetailsAtAdmin{}, errors.New(msg.ErrDataZero)
	}

	ok, err := au.adminRepository.CheckJobseekerById(id)

	if err != nil {
		return req.JobseekerDetailsAtAdmin{}, err
	}

	if !ok {
		return req.JobseekerDetailsAtAdmin{}, errors.New(msg.ErrIdExist)
	}

	data, err := au.adminRepository.GetJobseekerDetails(id)
	if err != nil {
		return req.JobseekerDetailsAtAdmin{}, err
	}
	return data, nil
}

func (au *adminUseCase) GetRecruiterDetails(id int) (req.RecruiterDetailsAtAdmin, error) {

	if id <= 0 {
		return req.RecruiterDetailsAtAdmin{}, errors.New(msg.ErrDataZero)
	}

	ok, err := au.adminRepository.CheckRecruiterById(id)

	if err != nil {
		return req.RecruiterDetailsAtAdmin{}, err
	}

	if !ok {
		return req.RecruiterDetailsAtAdmin{}, errors.New(msg.ErrIdExist)
	}

	data, err := au.adminRepository.GetRecruiterDetails(id)

	if err != nil {
		return req.RecruiterDetailsAtAdmin{}, err
	}
	return data, nil
}

// policies
func (au *adminUseCase) CreatePolicy(data req.CreatePolicyReq) (req.CreatePolicyRes, error) {

	if data.Title == "" {
		return req.CreatePolicyRes{}, errors.New("title not be null")
	}
	if data.Content == "" {
		return req.CreatePolicyRes{}, errors.New("content not be null")
	}
	pData, err := au.adminRepository.CreatePolicy(data)
	if err != nil {
		return req.CreatePolicyRes{}, err
	}

	return pData, nil
}

func (au *adminUseCase) UpdatePolicy(data req.UpdatePolicyReq) (req.CreatePolicyRes, error) {
	if data.Id <= 0 {
		return req.CreatePolicyRes{}, errors.New(msg.ErrDataZero)
	}
	if data.Title == "" {
		return req.CreatePolicyRes{}, errors.New("title not be null")
	}
	if data.Content == "" {
		return req.CreatePolicyRes{}, errors.New("content not be null")
	}
	ok, err := au.adminRepository.IsPolicyExist(data.Id)
	if err != nil {
		return req.CreatePolicyRes{}, err
	}
	if !ok {
		return req.CreatePolicyRes{}, errors.New(msg.ErrIdExist)
	}
	pData, err := au.adminRepository.UpdatePolicy(data)
	if err != nil {
		return req.CreatePolicyRes{}, err
	}
	return pData, nil
}

func (au *adminUseCase) DeletePolicy(policy_id int) (bool, error) {
	if policy_id <= 0 {
		return false, errors.New(msg.ErrDataZero)
	}
	ok, err := au.adminRepository.IsPolicyExist(policy_id)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New(msg.ErrIdExist)
	}
	okP, err := au.adminRepository.DeletePolicy(policy_id)
	if err != nil {
		return false, err
	}
	return okP, nil
}

func (au *adminUseCase) GetAllPolicies() (req.GetAllPolicyRes, error) {
	data, err := au.adminRepository.GetAllPolicies()
	if err != nil {
		return req.GetAllPolicyRes{}, err
	}
	return data, nil
}

func (au *adminUseCase) GetOnePolicy(policy_id int) (req.CreatePolicyRes, error) {
	if policy_id <= 0 {
		return req.CreatePolicyRes{}, errors.New(msg.ErrDataZero)
	}
	ok, err := au.adminRepository.IsPolicyExist(policy_id)
	if err != nil {
		return req.CreatePolicyRes{}, err
	}
	if !ok {
		return req.CreatePolicyRes{}, errors.New(msg.ErrIdExist)
	}
	data, err := au.adminRepository.GetOnePolicy(policy_id)
	if err != nil {
		return req.CreatePolicyRes{}, err
	}
	return data, nil
}
