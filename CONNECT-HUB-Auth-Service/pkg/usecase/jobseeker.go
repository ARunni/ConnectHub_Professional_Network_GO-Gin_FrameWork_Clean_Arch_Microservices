package usecase

import (
	"ConnetHub_auth/pkg/helper"
	repo "ConnetHub_auth/pkg/repository/interface"
	usecase "ConnetHub_auth/pkg/usecase/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"errors"

	msg "github.com/ARunni/Error_Message"
)

type jobseekerUseCase struct {
	jobseekerRepository repo.JobseekerRepository
}

func NewJobseekerUseCase(repo repo.JobseekerRepository) usecase.JobSeekerUseCase {
	return &jobseekerUseCase{
		jobseekerRepository: repo,
	}
}

func (ju *jobseekerUseCase) JobSeekerSignup(jobseekerdata req.JobSeekerSignUp) (req.TokenJobSeeker, error) {

	// Validations
	if jobseekerdata.Email == "" {
		return req.TokenJobSeeker{}, errors.New("email " + msg.ErrFieldEmpty)
	}
	if jobseekerdata.Password == "" {
		return req.TokenJobSeeker{}, errors.New("password " + msg.ErrFieldEmpty)
	}
	if jobseekerdata.FirstName == "" {
		return req.TokenJobSeeker{}, errors.New("first_name " + msg.ErrFieldEmpty)
	}
	if jobseekerdata.Gender == "" {
		return req.TokenJobSeeker{}, errors.New("gender " + msg.ErrFieldEmpty)
	}
	if jobseekerdata.PhoneNumber == "" {
		return req.TokenJobSeeker{}, errors.New("phone_number " + msg.ErrFieldEmpty)
	}
	if jobseekerdata.DateOfBirth == "" {
		return req.TokenJobSeeker{}, errors.New("date_of_birth " + msg.ErrFieldEmpty)
	}

	ok, err := ju.jobseekerRepository.CheckJobseekerExistsByEmail(jobseekerdata.Email)
	if err != nil {
		return req.TokenJobSeeker{}, err
	}
	if ok {
		return req.TokenJobSeeker{}, errors.New(msg.ErrAlreadyUser)
	}
	if jobseekerdata.Password != jobseekerdata.ConfirmPassword {
		return req.TokenJobSeeker{}, errors.New(msg.ErrPasswordMatch)
	}
	// Hashing Password

	hashedPassword, err := helper.PasswordHash(jobseekerdata.Password)
	if err != nil {
		return req.TokenJobSeeker{}, err
	}
	jobseekerdata.Password = hashedPassword

	// Inserting Data
	tokenResp, err := ju.jobseekerRepository.JobSeekerSignup(jobseekerdata)
	if err != nil {
		return req.TokenJobSeeker{}, err
	}
	// Generating Token
	token, err := helper.GenerateTokenJobseeker(tokenResp)
	if err != nil {
		return req.TokenJobSeeker{}, err
	}

	return req.TokenJobSeeker{
		JobSeeker: tokenResp,
		Token:     token,
	}, nil
}

func (ju *jobseekerUseCase) JobSeekerLogin(jobseekerDetails req.JobSeekerLogin) (req.TokenJobSeeker, error) {
	//  Validation
	if jobseekerDetails.Email == "" {
		return req.TokenJobSeeker{}, errors.New("email " + msg.ErrFieldEmpty)
	}
	if jobseekerDetails.Password == "" {
		return req.TokenJobSeeker{}, errors.New("password " + msg.ErrFieldEmpty)
	}
	ok, err := ju.jobseekerRepository.CheckJobseekerExistsByEmail(jobseekerDetails.Email)
	if err != nil {
		return req.TokenJobSeeker{}, err
	}
	if !ok {
		return req.TokenJobSeeker{}, errors.New(msg.ErrUserExistFalse)
	}
	okk, err := ju.jobseekerRepository.CheckJobseekerBlockByEmail(jobseekerDetails.Email)

	if err != nil {
		return req.TokenJobSeeker{}, err
	}
	if okk {
		return req.TokenJobSeeker{}, errors.New(msg.ErrUserBlockTrue)
	}

	jobseekerCompare, err := ju.jobseekerRepository.JobseekerLogin(jobseekerDetails)
	if err != nil {
		return req.TokenJobSeeker{}, err
	}

	// Comparing Password
	err = helper.CompareHashAndPassword(jobseekerCompare.Password, jobseekerDetails.Password)
	if err != nil {
		return req.TokenJobSeeker{}, errors.New(msg.ErrPasswordMatch)
	}

	access, err := helper.GenerateTokenJobseeker(jobseekerCompare)
	if err != nil {
		return req.TokenJobSeeker{}, err
	}
	return req.TokenJobSeeker{
		JobSeeker: jobseekerCompare,
		Token:     access,
	}, nil

}
