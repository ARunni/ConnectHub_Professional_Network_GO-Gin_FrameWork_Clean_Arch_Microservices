package service

import (
	pb "ConnetHub_auth/pkg/pb/auth/recruiter"
	interfaces "ConnetHub_auth/pkg/usecase/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"context"
)

type RecruiterServer struct {
	recruiterUseCase interfaces.RecruiterUseCase
	pb.UnimplementedRecruiterServer
}

func NewRecruiterServer(useCase interfaces.RecruiterUseCase) pb.RecruiterServer {

	return &RecruiterServer{
		recruiterUseCase: useCase,
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
