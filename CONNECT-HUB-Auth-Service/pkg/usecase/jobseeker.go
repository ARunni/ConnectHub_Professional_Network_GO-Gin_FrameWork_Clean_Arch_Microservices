package usecase

import (
	logging "github.com/ARunni/ConnetHub_auth/Logging"
	"github.com/ARunni/ConnetHub_auth/pkg/helper"
	repo "github.com/ARunni/ConnetHub_auth/pkg/repository/interface"
	usecase "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"
	req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"
	"errors"
	"os"

	msg "github.com/ARunni/Error_Message"
	"github.com/sirupsen/logrus"
)

type jobseekerUseCase struct {
	jobseekerRepository repo.JobseekerRepository
	Logger              *logrus.Logger
	LogFile             *os.File
}

func NewJobseekerUseCase(repo repo.JobseekerRepository) usecase.JobSeekerUseCase {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &jobseekerUseCase{
		jobseekerRepository: repo,
		Logger:              logger,
		LogFile:             logFile,
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

	ok, err := ju.jobseekerRepository.CheckJobseekerBlockByEmail(jobseekerDetails.Email)
	if err != nil {
		return req.TokenJobSeeker{}, err
	}
	if ok {
		return req.TokenJobSeeker{}, errors.New(msg.ErrUserBlockTrue)
	}
	//  Validation
	if jobseekerDetails.Email == "" {
		return req.TokenJobSeeker{}, errors.New("email " + msg.ErrFieldEmpty)
	}
	if jobseekerDetails.Password == "" {
		return req.TokenJobSeeker{}, errors.New("password " + msg.ErrFieldEmpty)
	}
	ok, err = ju.jobseekerRepository.CheckJobseekerExistsByEmail(jobseekerDetails.Email)
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

func (ju *jobseekerUseCase) JobSeekerGetProfile(id int) (req.JobSeekerProfile, error) {
	if id <= 0 {
		return req.JobSeekerProfile{}, errors.New("id error")
	}
	ok, err := ju.jobseekerRepository.IsJobseekerBlocked(id)
	if err != nil {
		return req.JobSeekerProfile{}, err
	}
	if ok {
		return req.JobSeekerProfile{}, errors.New(msg.ErrUserBlockTrue)
	}

	data, err := ju.jobseekerRepository.JobSeekerGetProfile(id)
	if err != nil {
		return req.JobSeekerProfile{}, err
	}
	return data, nil
}

func (ju *jobseekerUseCase) JobSeekerEditProfile(data req.JobSeekerProfile) (req.JobSeekerProfile, error) {

	if data.ID <= 0 {
		return req.JobSeekerProfile{}, errors.New("id error")
	}

	if data.Email == "" {
		return req.JobSeekerProfile{}, errors.New("email " + msg.ErrFieldEmpty)
	}

	if data.FirstName == "" {
		return req.JobSeekerProfile{}, errors.New("first_name " + msg.ErrFieldEmpty)
	}
	if data.Gender == "" {
		return req.JobSeekerProfile{}, errors.New("gender " + msg.ErrFieldEmpty)
	}
	if data.PhoneNumber == "" {
		return req.JobSeekerProfile{}, errors.New("phone_number " + msg.ErrFieldEmpty)
	}
	if data.DateOfBirth == "" {
		return req.JobSeekerProfile{}, errors.New("date_of_birth " + msg.ErrFieldEmpty)

	}
	ok, err := ju.jobseekerRepository.IsJobseekerBlocked(int(data.ID))
	if err != nil {
		return req.JobSeekerProfile{}, err
	}
	if ok {
		return req.JobSeekerProfile{}, errors.New(msg.ErrUserBlockTrue)
	}

	data, err = ju.jobseekerRepository.JobSeekerEditProfile(data)
	if err != nil {
		return req.JobSeekerProfile{}, err
	}
	return data, nil
}

// policies
func (ju *jobseekerUseCase) GetAllPolicies() (req.GetAllPolicyRes, error) {
	data, err := ju.jobseekerRepository.GetAllPolicies()
	if err != nil {
		return req.GetAllPolicyRes{}, err
	}
	return data, nil
}

func (ju *jobseekerUseCase) GetOnePolicy(policy_id int) (req.CreatePolicyRes, error) {
	if policy_id <= 0 {
		return req.CreatePolicyRes{}, errors.New(msg.ErrDataZero)
	}
	ok, err := ju.jobseekerRepository.IsPolicyExist(policy_id)
	if err != nil {
		return req.CreatePolicyRes{}, err
	}
	if !ok {
		return req.CreatePolicyRes{}, errors.New(msg.ErrIdExist)
	}
	data, err := ju.jobseekerRepository.GetOnePolicy(policy_id)
	if err != nil {
		return req.CreatePolicyRes{}, err
	}
	return data, nil
}
