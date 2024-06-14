package client

import (
	logging "github.com/ARunni/connectHub_gateway/Logging"
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/auth/interface"
	"github.com/ARunni/connectHub_gateway/pkg/config"
	"github.com/ARunni/connectHub_gateway/pkg/helper"
	pb "github.com/ARunni/connectHub_gateway/pkg/pb/auth/jobseeker"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type jobseekerClient struct {
	Client  pb.JobseekerClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewJobSeekerAuthClient(cfg config.Config) interfaces.JobSeekerAuthClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	grpcConnection, err := grpc.Dial(cfg.ConnetHubAuth, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth jobseeker", err)
	}

	grpcClient := pb.NewJobseekerClient(grpcConnection)
	return &jobseekerClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}

}

func (jc *jobseekerClient) JobSeekerSignup(jobseekerData models.JobSeekerSignUp) (models.TokenJobSeeker, error) {
	jc.Logger.Info("JobSeekerSignup at client started")
	jobseeker, err := jc.Client.JobSeekerSignup(context.Background(), &pb.JobSeekerSignupRequest{
		Firstname:       jobseekerData.FirstName,
		Lastname:        jobseekerData.LastName,
		Password:        jobseekerData.Password,
		Email:           jobseekerData.Email,
		PhoneNumber:     jobseekerData.PhoneNumber,
		DateOfBirth:     jobseekerData.DateOfBirth,
		Gender:          jobseekerData.Gender,
		ConfirmPassword: jobseekerData.ConfirmPassword,
	})
	if err != nil {
		jc.Logger.Error("Error in JobSeekerSignup at client:", err)
		return models.TokenJobSeeker{}, err
	}
	jc.Logger.Info("JobSeekerSignup at client success")
	return models.TokenJobSeeker{
		JobSeeker: models.JobSeekerDetailsResponse{
			ID:          uint(jobseeker.JobSeekerDetails.Id),
			Email:       jobseeker.JobSeekerDetails.Email,
			FirstName:   jobseeker.JobSeekerDetails.Firstname,
			LastName:    jobseeker.JobSeekerDetails.Lastname,
			PhoneNumber: jobseeker.JobSeekerDetails.PhoneNumber,
			DateOfBirth: jobseeker.JobSeekerDetails.DateOfBirth,
		},
		Token: jobseeker.Token,
	}, nil
}

func (jc *jobseekerClient) JobSeekerLogin(jobseekerData models.JobSeekerLogin) (models.TokenJobSeeker, error) {
	jc.Logger.Info("JobSeekerLogin at client started")
	jobseeker, err := jc.Client.JobSeekerLogin(context.Background(), &pb.JobSeekerLoginRequest{
		Email:    jobseekerData.Email,
		Password: jobseekerData.Password,
	})
	if err != nil {
		jc.Logger.Error("Error in JobSeekerLogin at client:", err)
		return models.TokenJobSeeker{}, err
	}
	jc.Logger.Info("JobSeekerLogin at client success")

	return models.TokenJobSeeker{
		JobSeeker: models.JobSeekerDetailsResponse{
			ID:          uint(jobseeker.JobSeekerDetails.Id),
			FirstName:   jobseeker.JobSeekerDetails.Firstname,
			LastName:    jobseeker.JobSeekerDetails.Lastname,
			Email:       jobseeker.JobSeekerDetails.Email,
			PhoneNumber: jobseeker.JobSeekerDetails.PhoneNumber,
			Gender:      jobseeker.JobSeekerDetails.Gender,
			DateOfBirth: jobseeker.JobSeekerDetails.DateOfBirth,
		},
		Token: jobseeker.Token,
	}, nil

}

func (jc *jobseekerClient) JobSeekerGetProfile(id int) (models.JobSeekerProfile, error) {
	jc.Logger.Info("JobSeekerGetProfile at client started")
	profile, err := jc.Client.JobSeekerGetProfile(context.Background(), &pb.GetProfileRequest{
		Id: int32(id),
	})
	if err != nil {
		jc.Logger.Error("Error in JobSeekerGetProfile at client:", err)
		return models.JobSeekerProfile{}, err
	}
	jc.Logger.Info("JobSeekerGetProfile at client success")

	return models.JobSeekerProfile{
		ID:          uint(profile.Profile.Id),
		Email:       profile.Profile.Email,
		FirstName:   profile.Profile.FirstName,
		LastName:    profile.Profile.LastName,
		PhoneNumber: profile.Profile.Phone,
		DateOfBirth: profile.Profile.DateOfBirth,
		Gender:      profile.Profile.Gender,
	}, nil

}

func (jc *jobseekerClient) JobSeekerEditProfile(profile models.JobSeekerProfile) (models.JobSeekerProfile, error) {
	jc.Logger.Info("JobSeekerEditProfile at client started")
	profileData, err := jc.Client.JobSeekerEditProfile(context.Background(), &pb.JobSeekerEditProfileRequest{
		Profile: &pb.JobSeekerProfile{
			Id:          uint64(profile.ID),
			FirstName:   profile.FirstName,
			Gender:      profile.Gender,
			Email:       profile.Email,
			DateOfBirth: profile.DateOfBirth,
			Phone:       profile.PhoneNumber,
			LastName:    profile.LastName,
		},
	})
	if err != nil {
		jc.Logger.Error("Error in JobSeekerEditProfile at client:", err)
		return models.JobSeekerProfile{}, err
	}
	jc.Logger.Info("JobSeekerEditProfile at client success")

	return models.JobSeekerProfile{
		ID:          uint(profileData.Profile.Id),
		FirstName:   profile.FirstName,
		Gender:      profile.Gender,
		Email:       profile.Email,
		DateOfBirth: profile.DateOfBirth,
		PhoneNumber: profile.PhoneNumber,
		LastName:    profile.LastName,
	}, nil

}

func (jc *jobseekerClient) GetAllPolicies() (models.GetAllPolicyRes, error) {
	jc.Logger.Info("GetAllPolicies at client started")
	Data, err := jc.Client.GetAllPolicies(context.Background(), &pb.GetAllPoliciesRequest{})
	if err != nil {
		jc.Logger.Error("Error in GetAllPolicies at client:", err)
		return models.GetAllPolicyRes{}, err
	}

	var policies []models.Policy
	for _, policy := range Data.Policies {
		policies = append(policies, models.Policy{
			ID:        uint(policy.Id),
			Title:     policy.Title,
			Content:   policy.Content,
			CreatedAt: helper.FromProtoTimestamp(policy.CreatedAt),
			UpdatedAt: helper.FromProtoTimestamp(policy.UpdatedAt),
		})

	}
	jc.Logger.Info("GetAllPolicies at client success")
	return models.GetAllPolicyRes{
		Policies: policies,
	}, nil

}

func (jc *jobseekerClient) GetOnePolicy(policy_id int) (models.CreatePolicyRes, error) {
	jc.Logger.Info("GetOnePolicy at client started")
	Data, err := jc.Client.GetOnePolicy(context.Background(), &pb.GetOnePolicyRequest{Id: int64(policy_id)})
	if err != nil {
		jc.Logger.Error("Error in GetOnePolicy at client:", err)
		return models.CreatePolicyRes{}, err
	}
	jc.Logger.Info("GetOnePolicy at client success")

	return models.CreatePolicyRes{
		Policies: models.Policy{
			ID:        uint(Data.Policy.Id),
			Title:     Data.Policy.Title,
			Content:   Data.Policy.Content,
			CreatedAt: helper.FromProtoTimestamp(Data.Policy.CreatedAt),
			UpdatedAt: helper.FromProtoTimestamp(Data.Policy.UpdatedAt),
		},
	}, nil

}
