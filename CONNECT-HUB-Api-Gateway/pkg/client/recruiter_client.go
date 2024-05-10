package client

import (
	interfaces "connectHub_gateway/pkg/client/interface"
	"connectHub_gateway/pkg/config"
	pb "connectHub_gateway/pkg/pb/auth/recruiter"
	"connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type recruiterClient struct {
	Client pb.RecruiterClient
}

func NewRecruiterClient(cfg config.Config) interfaces.RecruiterClient {
	grpcConnection, err := grpc.Dial(cfg.ConnetHubAuth, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth recruiter", err)
	}
	grpcClient := pb.NewRecruiterClient(grpcConnection)
	return &recruiterClient{
		Client: grpcClient,
	}

}

func (rc *recruiterClient) RecruiterSignup(recruiterData models.RecruiterSignUp) (models.TokenRecruiter, error) {
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
		return models.TokenRecruiter{}, err
	}
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
	recruiter, err := rc.Client.RecruiterLogin(context.Background(), &pb.RecruiterLoginInRequest{
		Email:    recruiterData.Email,
		Password: recruiterData.Password,
	})
	if err != nil {
		return models.TokenRecruiter{}, err
	}
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
