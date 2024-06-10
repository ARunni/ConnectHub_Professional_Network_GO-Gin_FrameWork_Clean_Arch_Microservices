package service

import (
	logging "github.com/ARunni/ConnetHub_auth/Logging"
	pb "github.com/ARunni/ConnetHub_auth/pkg/pb/auth/recruiter"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"
	req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RecruiterServer struct {
	recruiterUseCase interfaces.RecruiterUseCase
	pb.UnimplementedRecruiterServer
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewRecruiterServer(useCase interfaces.RecruiterUseCase) pb.RecruiterServer {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &RecruiterServer{
		recruiterUseCase: useCase,
		Logger:           logger,
		LogFile:          logFile,
	}
}

func (rs *RecruiterServer) RecruiterSignup(ctx context.Context, Req *pb.RecruiterSignupRequest) (*pb.RecruiterSignupResponse, error) {
	recruiterSignup := req.RecruiterSignUp{
		Company_name:         Req.CompanyName,
		Industry:             Req.Industry,
		Company_size:         int(Req.CompanySize),
		Website:              Req.Website,
		Headquarters_address: Req.HeadquartersAddress,
		About_company:        Req.AboutCompany,
		Contact_email:        Req.Email,
		Contact_phone_number: uint(Req.PhoneNumber),
		Password:             Req.Password,
		ConfirmPassword:      Req.ConfirmPassword,
	}

	recruiter, err := rs.recruiterUseCase.RecruiterSignup(recruiterSignup)
	if err != nil {
		return nil, err
	}
	return &pb.RecruiterSignupResponse{
		Status: 200,
		RecruiterDetails: &pb.RecruiterDetails{
			Id:                  uint64(recruiter.Recruiter.ID),
			CompanyName:         recruiter.Recruiter.Company_name,
			Email:               recruiter.Recruiter.Contact_email,
			AboutCompany:        recruiter.Recruiter.About_company,
			Industry:            recruiter.Recruiter.Industry,
			CompanySize:         int64(recruiter.Recruiter.Company_size),
			Website:             recruiter.Recruiter.Website,
			PhoneNumber:         int64(recruiter.Recruiter.Contact_phone_number),
			HeadquartersAddress: recruiter.Recruiter.Headquarters_address,
		},
		Token: recruiter.Token,
	}, nil
}

func (rs *RecruiterServer) RecruiterLogin(ctx context.Context, Req *pb.RecruiterLoginInRequest) (*pb.RecruiterLoginResponse, error) {
	recruiterLogin := req.RecruiterLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	recruiterdata, err := rs.recruiterUseCase.RecruiterLogin(recruiterLogin)
	if err != nil {
		return nil, err
	}
	return &pb.RecruiterLoginResponse{
		Status: 200,
		RecruiterDetails: &pb.RecruiterDetails{
			Id:                  uint64(recruiterdata.Recruiter.ID),
			CompanyName:         recruiterdata.Recruiter.Company_name,
			Email:               recruiterdata.Recruiter.Contact_email,
			AboutCompany:        recruiterdata.Recruiter.About_company,
			Industry:            recruiterdata.Recruiter.Industry,
			CompanySize:         int64(recruiterdata.Recruiter.Company_size),
			Website:             recruiterdata.Recruiter.Website,
			PhoneNumber:         int64(recruiterdata.Recruiter.Contact_phone_number),
			HeadquartersAddress: recruiterdata.Recruiter.Headquarters_address,
		},
		Token: recruiterdata.Token,
	}, nil

}

func (rs *RecruiterServer) RecruiterGetProfile(ctx context.Context, Req *pb.GetProfileRequest) (*pb.RecruiterDetailsResponse, error) {
	recruiterId := Req.RecruiterId
	recruiterdata, err := rs.recruiterUseCase.RecruiterGetProfile(int(recruiterId))
	if err != nil {
		return nil, err
	}
	return &pb.RecruiterDetailsResponse{
		Id:                  uint64(recruiterdata.ID),
		CompanyName:         recruiterdata.Company_name,
		Email:               recruiterdata.Contact_email,
		AboutCompany:        recruiterdata.About_company,
		Industry:            recruiterdata.Industry,
		CompanySize:         int64(recruiterdata.Company_size),
		Website:             recruiterdata.Website,
		PhoneNumber:         int64(recruiterdata.Contact_phone_number),
		HeadquartersAddress: recruiterdata.Headquarters_address,
	}, nil

}

func (rs *RecruiterServer) RecruiterEditProfile(ctx context.Context, Req *pb.RecruiterEditProfileRequest) (*pb.RecruiterEditProfileResponse, error) {
	recruiterProfile := req.RecruiterProfile{
		ID:                   uint(Req.Profile.Id),
		Company_name:         Req.Profile.CompanyName,
		Industry:             Req.Profile.Industry,
		Website:              Req.Profile.Website,
		Headquarters_address: Req.Profile.HeadquartersAddress,
		About_company:        Req.Profile.AboutCompany,
		Company_size:         int(Req.Profile.CompanySize),
		Contact_email:        Req.Profile.Email,
		Contact_phone_number: uint(Req.Profile.PhoneNumber),
	}
	recruiterdata, err := rs.recruiterUseCase.RecruiterEditProfile(recruiterProfile)
	if err != nil {
		return nil, err
	}
	return &pb.RecruiterEditProfileResponse{
		Profile: &pb.RecruiterDetails{
			Id:                  uint64(recruiterdata.ID),
			CompanyName:         recruiterdata.Company_name,
			Email:               recruiterdata.Contact_email,
			AboutCompany:        recruiterdata.About_company,
			Industry:            recruiterdata.Industry,
			CompanySize:         int64(recruiterdata.Company_size),
			Website:             recruiterdata.Website,
			PhoneNumber:         int64(recruiterdata.Contact_phone_number),
			HeadquartersAddress: recruiterdata.Headquarters_address,
		},
	}, nil

}

func (rs *RecruiterServer) GetAllPolicies(ctx context.Context, Req *pb.GetAllPoliciesRequest) (*pb.GetAllPoliciesResponse, error) {

	recruiterdata, err := rs.recruiterUseCase.GetAllPolicies()
	if err != nil {
		return nil, err
	}
	var policies []*pb.Policy

	for _, p := range recruiterdata.Policies {
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

func (rs *RecruiterServer) GetOnePolicy(ctx context.Context, Req *pb.GetOnePolicyRequest) (*pb.GetOnePolicyResponse, error) {
	policy_id := Req.Id
	recruiterdata, err := rs.recruiterUseCase.GetOnePolicy(int(policy_id))
	if err != nil {
		return nil, err
	}
	return &pb.GetOnePolicyResponse{
		Policy: &pb.Policy{
			Id:        int64(recruiterdata.Policies.ID),
			Title:     recruiterdata.Policies.Title,
			Content:   recruiterdata.Policies.Content,
			CreatedAt: timestamppb.New(recruiterdata.Policies.CreatedAt),
			UpdatedAt: timestamppb.New(recruiterdata.Policies.UpdatedAt),
		},
	}, nil

}
