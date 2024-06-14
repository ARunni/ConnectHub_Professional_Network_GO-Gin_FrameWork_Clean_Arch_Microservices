package client

import (
	"context"
	"fmt"
	"os"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/auth/interface"
	"github.com/ARunni/connectHub_gateway/pkg/config"
	"github.com/ARunni/connectHub_gateway/pkg/helper"
	pb "github.com/ARunni/connectHub_gateway/pkg/pb/auth/recruiter"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type recruiterClient struct {
	Client  pb.RecruiterClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewRecruiterAuthClient(cfg config.Config) interfaces.RecruiterAuthClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	grpcConnection, err := grpc.Dial(cfg.ConnetHubAuth, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth recruiter", err)
	}
	grpcClient := pb.NewRecruiterClient(grpcConnection)
	return &recruiterClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}

}

func (rc *recruiterClient) RecruiterSignup(recruiterData models.RecruiterSignUp) (models.TokenRecruiter, error) {
	rc.Logger.Info("RecruiterSignup at client started")
	recruiter, err := rc.Client.RecruiterSignup(context.Background(), &pb.RecruiterSignupRequest{
		CompanyName:         recruiterData.Company_name,
		Email:               recruiterData.Contact_email,
		Password:            recruiterData.Password,
		AboutCompany:        recruiterData.About_company,
		Industry:            recruiterData.Industry,
		CompanySize:         int64(recruiterData.Company_size),
		Website:             recruiterData.Website,
		PhoneNumber:         int64(recruiterData.Contact_phone_number),
		HeadquartersAddress: recruiterData.Headquarters_address,
		ConfirmPassword:     recruiterData.ConfirmPassword,
	})
	if err != nil {
		rc.Logger.Error("Error in RecruiterSignup at client: ", err)
		return models.TokenRecruiter{}, err
	}
	rc.Logger.Info("RecruiterSignup at client success")
	return models.TokenRecruiter{
		Recruiter: models.RecruiterDetailsResponse{
			ID:                   uint(recruiter.RecruiterDetails.Id),
			Company_name:         recruiter.RecruiterDetails.CompanyName,
			Industry:             recruiter.RecruiterDetails.Industry,
			Company_size:         int(recruiter.RecruiterDetails.CompanySize),
			Website:              recruiter.RecruiterDetails.Website,
			Headquarters_address: recruiter.RecruiterDetails.HeadquartersAddress,
			About_company:        recruiter.RecruiterDetails.AboutCompany,
			Contact_email:        recruiter.RecruiterDetails.Email,
			Contact_phone_number: uint(recruiter.RecruiterDetails.PhoneNumber),
		},
		Token: recruiter.Token,
	}, nil
}

func (rc *recruiterClient) RecruiterLogin(recruiterData models.RecruiterLogin) (models.TokenRecruiter, error) {
	rc.Logger.Info("RecruiterLogin at client started")
	recruiter, err := rc.Client.RecruiterLogin(context.Background(), &pb.RecruiterLoginInRequest{
		Email:    recruiterData.Email,
		Password: recruiterData.Password,
	})
	if err != nil {
		rc.Logger.Error("Error in RecruiterLogin at client: ", err)
		return models.TokenRecruiter{}, err
	}
	rc.Logger.Info("RecruiterLogin at client success")
	return models.TokenRecruiter{
		Recruiter: models.RecruiterDetailsResponse{
			ID:                   uint(recruiter.RecruiterDetails.Id),
			Company_name:         recruiter.RecruiterDetails.CompanyName,
			Industry:             recruiter.RecruiterDetails.Industry,
			Company_size:         int(recruiter.RecruiterDetails.CompanySize),
			Website:              recruiter.RecruiterDetails.Website,
			Headquarters_address: recruiter.RecruiterDetails.HeadquartersAddress,
			About_company:        recruiter.RecruiterDetails.AboutCompany,
			Contact_email:        recruiter.RecruiterDetails.Email,
			Contact_phone_number: uint(recruiter.RecruiterDetails.PhoneNumber),
		},
		Token: recruiter.Token,
	}, nil
}

func (rc *recruiterClient) RecruiterGetProfile(id int) (models.RecruiterProfile, error) {
	rc.Logger.Info("RecruiterGetProfile at client started")
	data, err := rc.Client.RecruiterGetProfile(context.Background(), &pb.GetProfileRequest{
		RecruiterId: int32(id),
	})
	if err != nil {
		rc.Logger.Error("Error in RecruiterGetProfile at client: ", err)
		return models.RecruiterProfile{}, err
	}
	rc.Logger.Info("RecruiterGetProfile at client success")
	return models.RecruiterProfile{
		ID:                   uint(data.Id),
		Company_name:         data.CompanyName,
		Industry:             data.Industry,
		Company_size:         int(data.CompanySize),
		Website:              data.Website,
		Headquarters_address: data.HeadquartersAddress,
		About_company:        data.AboutCompany,
		Contact_email:        data.Email,
		Contact_phone_number: uint(data.PhoneNumber),
	}, nil
}

func (rc *recruiterClient) RecruiterEditProfile(data models.RecruiterProfile) (models.RecruiterProfile, error) {
	rc.Logger.Info("RecruiterEditProfile at client started")
	profiledata, err := rc.Client.RecruiterEditProfile(context.Background(), &pb.RecruiterEditProfileRequest{
		Profile: &pb.RecruiterDetails{
			Id:                  uint64(data.ID),
			CompanyName:         data.Company_name,
			Email:               data.Contact_email,
			AboutCompany:        data.About_company,
			Industry:            data.Industry,
			CompanySize:         int64(data.Company_size),
			Website:             data.Website,
			PhoneNumber:         int64(data.Contact_phone_number),
			HeadquartersAddress: data.Headquarters_address,
		},
	})
	if err != nil {
		rc.Logger.Error("Error in RecruiterEditProfile at client: ", err)
		return models.RecruiterProfile{}, err
	}
	rc.Logger.Info("RecruiterEditProfile at client success")
	return models.RecruiterProfile{
		ID:                   uint(profiledata.Profile.Id),
		Company_name:         profiledata.Profile.CompanyName,
		Industry:             profiledata.Profile.Industry,
		Company_size:         int(profiledata.Profile.CompanySize),
		Website:              profiledata.Profile.Website,
		Headquarters_address: profiledata.Profile.HeadquartersAddress,
		About_company:        profiledata.Profile.AboutCompany,
		Contact_email:        profiledata.Profile.Email,
		Contact_phone_number: uint(profiledata.Profile.PhoneNumber),
	}, nil
}

func (rc *recruiterClient) GetAllPolicies() (models.GetAllPolicyRes, error) {
	rc.Logger.Info("GetAllPolicies at client started")
	data, err := rc.Client.GetAllPolicies(context.Background(), &pb.GetAllPoliciesRequest{})
	if err != nil {
		rc.Logger.Error("Error in GetAllPolicies at client: ", err)
		return models.GetAllPolicyRes{}, err
	}
	var policies []models.Policy
	for _, policy := range data.Policies {
		policies = append(policies, models.Policy{
			ID:        uint(policy.Id),
			Title:     policy.Title,
			Content:   policy.Content,
			CreatedAt: helper.FromProtoTimestamp(policy.CreatedAt),
			UpdatedAt: helper.FromProtoTimestamp(policy.UpdatedAt),
		})

	}
	rc.Logger.Info("GetAllPolicies at client success")
	return models.GetAllPolicyRes{
		Policies: policies,
	}, nil
}

func (rc *recruiterClient) GetOnePolicy(policy_id int) (models.CreatePolicyRes, error) {
	rc.Logger.Info("GetOnePolicy at client started")
	data, err := rc.Client.GetOnePolicy(context.Background(), &pb.GetOnePolicyRequest{Id: int64(policy_id)})
	if err != nil {
		rc.Logger.Error("Error in GetOnePolicy at client: ", err)
		return models.CreatePolicyRes{}, err
	}
	rc.Logger.Info("GetOnePolicy at client success")

	return models.CreatePolicyRes{
		Policies: models.Policy{
			ID:        uint(data.Policy.Id),
			Title:     data.Policy.Title,
			Content:   data.Policy.Content,
			CreatedAt: helper.FromProtoTimestamp(data.Policy.CreatedAt),
			UpdatedAt: helper.FromProtoTimestamp(data.Policy.UpdatedAt),
		},
	}, nil
}
