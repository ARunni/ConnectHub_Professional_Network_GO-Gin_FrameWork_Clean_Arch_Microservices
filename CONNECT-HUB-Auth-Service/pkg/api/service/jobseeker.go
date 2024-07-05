package service

import (
	"context"
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	pb "github.com/ARunni/ConnetHub_auth/pkg/pb/auth/jobseeker"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"
	req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"

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
	js.Logger.Info("JobSeekerSignup at JobSeekerServer started")
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

	js.Logger.Info("JobSeekerSignup at jobseekerUsecase started")
	jobseekerData, err := js.jobseekerUsecase.JobSeekerSignup(jobseekerSignup)
	if err != nil {
		js.Logger.Error("Error at jobseekerUsecase", err)
		return nil, err
	}
	js.Logger.Info("JobSeekerSignup at jobseekerUsecase success")
	js.Logger.Info("JobSeekerSignup at JobSeekerServer success")
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
	js.Logger.Info("JobSeekerLogin at JobSeekerServer started")
	jobseekerLogin := req.JobSeekerLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	js.Logger.Info("JobSeekerLogin at jobseekerUsecase started")
	jobseekerData, err := js.jobseekerUsecase.JobSeekerLogin(jobseekerLogin)
	if err != nil {
		js.Logger.Error("Error at jobseekerUsecase", err)
		return nil, err
	}
	js.Logger.Info("JobSeekerLogin at jobseekerUsecase success")
	js.Logger.Info("JobSeekerLogin at JobSeekerServer success")
	return &pb.JobSeekerLoginResponse{
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

func (js *JobSeekerServer) JobSeekerGetProfile(ctx context.Context, Req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	js.Logger.Info("JobSeekerGetProfile at JobSeekerServer started")
	jobseekerId := Req.Id
	js.Logger.Info("JobSeekerGetProfile at jobseekerUsecase started")
	jobseekerData, err := js.jobseekerUsecase.JobSeekerGetProfile(int(jobseekerId))
	if err != nil {
		js.Logger.Error("Error at jobseekerUsecase", err)
		return nil, err
	}
	js.Logger.Info("JobSeekerGetProfile at jobseekerUsecase success")
	js.Logger.Info("JobSeekerGetProfile at JobSeekerServer success")
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
	js.Logger.Info("JobSeekerGetProfile at JobSeekerServer started")
	jobseekerReq := req.JobSeekerProfile{
		ID:          uint(Req.Profile.Id),
		FirstName:   Req.Profile.FirstName,
		Gender:      Req.Profile.Gender,
		Email:       Req.Profile.Email,
		LastName:    Req.Profile.LastName,
		PhoneNumber: Req.Profile.Phone,
		DateOfBirth: Req.Profile.DateOfBirth,
	}
	js.Logger.Info("JobSeekerEditProfile at jobseekerUsecase started")

	jobseekerData, err := js.jobseekerUsecase.JobSeekerEditProfile(jobseekerReq)
	if err != nil {
		js.Logger.Error("Error at jobseekerUsecase", err)
		return nil, err
	}
	js.Logger.Info("JobSeekerEditProfile at jobseekerUsecase success")
	js.Logger.Info("JobSeekerEditProfile at JobSeekerServer success")
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
	js.Logger.Info("GetAllPolicies at JobSeekerServer started")

	js.Logger.Info("GetAllPolicies at jobseekerUsecase started")
	jobseekerData, err := js.jobseekerUsecase.GetAllPolicies()
	if err != nil {
		js.Logger.Error("Error at jobseekerUsecase", err)
		return nil, err
	}
	js.Logger.Info("GetAllPolicies at jobseekerUsecase success")
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

	js.Logger.Info("GetAllPolicies at JobSeekerServer success")

	return &pb.GetAllPoliciesResponse{
		Policies: policies,
	}, nil
}

func (js *JobSeekerServer) GetOnePolicy(ctx context.Context, Req *pb.GetOnePolicyRequest) (*pb.GetOnePolicyResponse, error) {

	js.Logger.Info("GetOnePolicy at JobSeekerServer started")

	policy_id := Req.Id
	js.Logger.Info("GetOnePolicy at jobseekerUsecase started")
	jobseekerData, err := js.jobseekerUsecase.GetOnePolicy(int(policy_id))
	if err != nil {
		js.Logger.Error("Error at jobseekerUsecase", err)
		return nil, err
	}
	js.Logger.Info("GetOnePolicy at jobseekerUsecase success")
	js.Logger.Info("GetOnePolicy at JobSeekerServer success")
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
