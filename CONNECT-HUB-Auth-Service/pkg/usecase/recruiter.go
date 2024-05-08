package usecase

import (
	"ConnetHub_auth/pkg/helper"
	repo "ConnetHub_auth/pkg/repository/interface"
	usecase "ConnetHub_auth/pkg/usecase/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"errors"
	"strconv"

	msg "github.com/ARunni/Error_Message"
)

type recruiterUseCase struct {
	recruiterRepository repo.RecruiterRepository
}

func NewRecruiterUseCase(repo repo.RecruiterRepository) usecase.RecruiterUseCase {
	return &recruiterUseCase{
		recruiterRepository: repo,
	}
}

func (ju *recruiterUseCase) RecruiterSignup(recruiterdata req.RecruiterSignUp) (req.TokenRecruiter, error) {
	// Validations
	if recruiterdata.Company_name == "" {
		return req.TokenRecruiter{}, errors.New(msg.ErrFieldEmpty)
	}
	if recruiterdata.About_company == "" {
		return req.TokenRecruiter{}, errors.New(msg.ErrFieldEmpty)
	}
	if recruiterdata.Company_size < 0 {
		return req.TokenRecruiter{}, errors.New(msg.ErrFieldEmpty)
	}
	if recruiterdata.Contact_email == "" {
		return req.TokenRecruiter{}, errors.New(msg.ErrFieldEmpty)
	}

	phoneStr := strconv.Itoa(int(recruiterdata.Contact_phone_number))

	if helper.ValidatePhoneNumber(phoneStr) {
		return req.TokenRecruiter{}, errors.New(msg.ErrInvalidPhone)
	}

	if recruiterdata.Password == "" {
		return req.TokenRecruiter{}, errors.New(msg.ErrFieldEmpty)
	}
	if recruiterdata.Password != recruiterdata.ConfirmPassword {
		return req.TokenRecruiter{}, errors.New(msg.ErrPasswordMatch)
	}

	ok, err := ju.recruiterRepository.CheckRecruiterExistsByEmail(recruiterdata.Contact_email)
	if err != nil {
		return req.TokenRecruiter{}, err
	}
	if ok {
		return req.TokenRecruiter{}, errors.New(msg.ErrAlreadyUser)
	}

	recruiterResp, err := ju.recruiterRepository.RecruiterSignup(recruiterdata)
	if err != nil {
		return req.TokenRecruiter{}, err
	}
	access, err := helper.GenerateTokenRecruiter(recruiterResp)
	if err != nil {
		return req.TokenRecruiter{}, err
	}
	return req.TokenRecruiter{
		Recruiter: recruiterResp,
		Token:     access,
	}, nil

}


func (ju *recruiterUseCase) RecruiterLogin(recruiterDetails req.RecruiterLogin) (req.TokenRecruiter, error){
	
}
