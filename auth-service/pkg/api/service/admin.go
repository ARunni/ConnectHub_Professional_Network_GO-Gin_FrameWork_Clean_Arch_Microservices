package service

import (
	pb "ConnetHub_auth/pkg/pb/auth/admin"
	interfaces "ConnetHub_auth/pkg/usecase/interface"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"context"
)

type AdminServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAdminServer
}

func NewAdminServer(useCase interfaces.AdminUseCase) pb.AdminServer {

	return &AdminServer{
		adminUseCase: useCase,
	}
}

func (ad *AdminServer) AdminLogin(ctx context.Context, Req *pb.AdminLoginInRequest) (*pb.AdminLoginResponse, error) {
	adminLogin := req.AdminLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	admin, err := ad.adminUseCase.LoginHandler(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(admin.Admin.ID),
		Firstname: admin.Admin.Firstname,
		Email:     admin.Admin.Email,
	}
	return &pb.AdminLoginResponse{
		Status:       200,
		AdminDetails: adminDetails,
		Token:        admin.Token,
	}, nil
}
