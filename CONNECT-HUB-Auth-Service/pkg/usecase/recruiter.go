package usecase

import (
	"ConnetHub_auth/pkg/helper"
	repo "ConnetHub_auth/pkg/repository/interface"
	usecase "ConnetHub_auth/pkg/usecase/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"errors"

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
		return req.TokenRecruiter{}, errors.New("company_name " + msg.ErrFieldEmpty)
	}
	if recruiterdata.About_company == "" {
		return req.TokenRecruiter{}, errors.New("about_company " + msg.ErrFieldEmpty)
	}
	if recruiterdata.Company_size < 0 {
		return req.TokenRecruiter{}, errors.New("company_size " + msg.ErrFieldEmpty)
	}
	if recruiterdata.Contact_email == "" {
		return req.TokenRecruiter{}, errors.New("contact_email " + msg.ErrFieldEmpty)
	}

	// phoneStr := strconv.Itoa(int(recruiterdata.Contact_phone_number))

	// if helper.ValidatePhoneNumber(phoneStr) {
	// 	return req.TokenRecruiter{}, errors.New(msg.ErrInvalidPhone)
	// }

	if recruiterdata.Password == "" {
		return req.TokenRecruiter{}, errors.New("password " + msg.ErrFieldEmpty)
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
	hashPassword, err := helper.PasswordHash(recruiterdata.Password)
	if err != nil {
		return req.TokenRecruiter{}, err
	}
	recruiterdata.Password = hashPassword

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

func (ju *recruiterUseCase) RecruiterLogin(recruiterDetails req.RecruiterLogin) (req.TokenRecruiter, error) {
	// validation
	if recruiterDetails.Email == "" {
		return req.TokenRecruiter{}, errors.New("email " + msg.ErrFieldEmpty)
	}
	if recruiterDetails.Password == "" {
		return req.TokenRecruiter{}, errors.New("password " + msg.ErrFieldEmpty)
	}
	ok, err := ju.recruiterRepository.CheckRecruiterExistsByEmail(recruiterDetails.Email)
	if err != nil {
		return req.TokenRecruiter{}, err
	}
	if !ok {
		return req.TokenRecruiter{}, errors.New(msg.ErrUserExistFalse)
	}

	okk, err := ju.recruiterRepository.CheckRecruiterBlockByEmail(recruiterDetails.Email)

	if err != nil {
		return req.TokenRecruiter{}, err
	}
	if okk {
		return req.TokenRecruiter{}, errors.New(msg.ErrUserBlockTrue)
	}

	recruiterCompare, err := ju.recruiterRepository.RecruiterLogin(recruiterDetails)
	if err != nil {
		return req.TokenRecruiter{}, err
	}

	// Comparing Password
	err = helper.CompareHashAndPassword(recruiterCompare.Password, recruiterDetails.Password)
	if err != nil {
		return req.TokenRecruiter{}, errors.New(msg.ErrPasswordMatch)
	}
	access, err := helper.GenerateTokenRecruiter(recruiterCompare)
	if err != nil {
		return req.TokenRecruiter{}, err
	}
	return req.TokenRecruiter{
		Recruiter: recruiterCompare,
		Token:     access,
	}, nil
}
