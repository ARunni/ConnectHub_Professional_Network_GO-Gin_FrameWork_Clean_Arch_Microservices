package service

import (
	logging "github.com/ARunni/ConnetHub_auth/Logging"
	pb "github.com/ARunni/ConnetHub_auth/pkg/pb/auth/jobseeker"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"
	req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type JobSeekerServer struct {
	jobseekerUsecase interfaces.JobSeekerUseCase
	pb.UnimplementedJobseekerServer
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewJobSeekerServer(useCase interfaces.JobSeekerUseCase) pb.JobseekerServer {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &JobSeekerServer{
		jobseekerUsecase: useCase,
		Logger:           logger,
		LogFile:          logFile,
	}
}

func (js *JobSeekerServer) JobSeekerSignup(ctx context.Context, Req *pb.JobSeekerSignupRequest) (*pb.JobSeekerSignupResponse, error) {
	jobseekerSignup := req.JobSeekerSignUp{
		Email:           Req.Email,
		Password:        Req.Password,
		ConfirmPassword: Req.ConfirmPassword,
		FirstName:       Req.Firstname,
		LastName:        Req.Lastname,
		PhoneNumber:     Req.PhoneNumber,
		DateOfBirth:     Req.DateOfBirth,
		Gender:          Req.Gender,
	}

	jobseekerData, err := js.jobseekerUsecase.JobSeekerSignup(jobseekerSignup)
	if err != nil {
		return nil, err
	}
	return &pb.JobSeekerSignupResponse{
		Status: 200,
		JobSeekerDetails: &pb.JobSeekerDetails{
			Id:          uint64(jobseekerData.JobSeeker.ID),
			Email:       jobseekerData.JobSeeker.Email,
			Firstname:   jobseekerData.JobSeeker.FirstName,
			Lastname:    jobseekerData.JobSeeker.LastName,
			PhoneNumber: jobseekerData.JobSeeker.PhoneNumber,
			DateOfBirth: jobseekerData.JobSeeker.DateOfBirth,
			Gender:      jobseekerData.JobSeeker.Gender,
		},
		Token: jobseekerData.Token,
	}, nil

}

func (js *JobSeekerServer) JobSeekerLogin(ctx context.Context, Req *pb.JobSeekerLoginRequest) (*pb.JobSeekerLoginResponse, error) {
	jobseekerLogin := req.JobSeekerLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	jobseekerData, err := js.jobseekerUsecase.JobSeekerLogin(jobseekerLogin)
	if err != nil {
		return nil, err
	}
	return &pb.JobSeekerLoginResponse{
		Status: 200,
		JobSeekerDetails: &pb.JobSeekerDetails{
			Id:          uint64(jobseekerData.JobSeeker.ID),
			Email:       jobseekerData.JobSeeker.Email,
			Firstname:   jobseekerData.JobSeeker.FirstName,
			Lastname:    jobseekerData.JobSeeker.LastName,
			PhoneNumber: jobseekerData.JobSeeker.PhoneNumber,
			DateOfBirth: jobseekerData.JobSeeker.DateOfBirth,
		},
		Token: jobseekerData.Token,
	}, nil
}

func (js *JobSeekerServer) JobSeekerGetProfile(ctx context.Context, Req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	jobseekerId := Req.Id
	jobseekerData, err := js.jobseekerUsecase.JobSeekerGetProfile(int(jobseekerId))
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileResponse{
		Status: 200,
		Profile: &pb.JobSeekerProfile{
			Id:          uint64(jobseekerData.ID),
			FirstName:   jobseekerData.FirstName,
			Gender:      jobseekerData.Gender,
			Email:       jobseekerData.Email,
			DateOfBirth: jobseekerData.DateOfBirth,
			Phone:       jobseekerData.PhoneNumber,
			LastName:    jobseekerData.LastName,
		},
	}, nil
}

func (js *JobSeekerServer) JobSeekerEditProfile(ctx context.Context, Req *pb.JobSeekerEditProfileRequest) (*pb.JobSeekerEditProfileResponse, error) {
	jobseekerReq := req.JobSeekerProfile{
		ID:          uint(Req.Profile.Id),
		FirstName:   Req.Profile.FirstName,
		Gender:      Req.Profile.Gender,
		Email:       Req.Profile.Email,
		LastName:    Req.Profile.LastName,
		PhoneNumber: Req.Profile.Phone,
		DateOfBirth: Req.Profile.DateOfBirth,
	}

	jobseekerData, err := js.jobseekerUsecase.JobSeekerEditProfile(jobseekerReq)
	if err != nil {
		return nil, err
	}
	return &pb.JobSeekerEditProfileResponse{

		Profile: &pb.JobSeekerProfile{
			Id:          uint64(jobseekerData.ID),
			FirstName:   jobseekerData.FirstName,
			Gender:      jobseekerData.Gender,
			Email:       jobseekerData.Email,
			DateOfBirth: jobseekerData.DateOfBirth,
			Phone:       jobseekerData.PhoneNumber,
			LastName:    jobseekerData.LastName,
		},
	}, nil
}

func (js *JobSeekerServer) GetAllPolicies(ctx context.Context, Req *pb.GetAllPoliciesRequest) (*pb.GetAllPoliciesResponse, error) {

	jobseekerData, err := js.jobseekerUsecase.GetAllPolicies()
	if err != nil {
		return nil, err
	}
	var policies []*pb.Policy

	for _, p := range jobseekerData.Policies {
		policy := &pb.Policy{
			Id:        int64(p.ID),
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: timestamppb.New(p.CreatedAt),
			UpdatedAt: timestamppb.New(p.UpdatedAt),
		}

		policies = append(policies, policy)
	}

	return &pb.GetAllPoliciesResponse{
		Policies: policies,
	}, nil
}

func (js *JobSeekerServer) GetOnePolicy(ctx context.Context, Req *pb.GetOnePolicyRequest) (*pb.GetOnePolicyResponse, error) {

	policy_id := Req.Id
	jobseekerData, err := js.jobseekerUsecase.GetOnePolicy(int(policy_id))
	if err != nil {
		return nil, err
	}
	return &pb.GetOnePolicyResponse{
		Policy: &pb.Policy{
			Id:        int64(jobseekerData.Policies.ID),
			Title:     jobseekerData.Policies.Title,
			Content:   jobseekerData.Policies.Content,
			CreatedAt: timestamppb.New(jobseekerData.Policies.CreatedAt),
			UpdatedAt: timestamppb.New(jobseekerData.Policies.UpdatedAt),
		},
	}, nil
}
