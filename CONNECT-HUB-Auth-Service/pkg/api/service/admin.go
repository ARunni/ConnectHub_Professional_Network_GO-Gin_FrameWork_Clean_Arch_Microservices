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
	admin, err := ad.adminUseCase.AdminLogin(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(admin.Admin.ID),
		Firstname: admin.Admin.Firstname,
		Lastname:  admin.Admin.Lastname,
		Email:     admin.Admin.Email,
	}
	return &pb.AdminLoginResponse{
		Status:       200,
		AdminDetails: adminDetails,
		Token:        admin.Token,
	}, nil
}

func (as *AdminServer) GetJobseekers(ctx context.Context, Req *pb.GetJobseekerRequest) (*pb.GetJobseekerResponse, error) {
	page := Req.Page
	GetJobseeker, err := as.adminUseCase.GetJobseekers(int(page))
	if err != nil {
		return &pb.GetJobseekerResponse{}, err
	}
	var jobseekerDetails []*pb.JobSeekerDetails
	for _, data := range GetJobseeker {
		jobseekerDetails = append(jobseekerDetails, &pb.JobSeekerDetails{
			Id:          uint64(data.Id),
			Email:       data.Email,
			Firstname:   data.Name,
			PhoneNumber: data.Phone,
			Blocked:     data.Blocked,
		})
	}
	return &pb.GetJobseekerResponse{
		Jobseeker: jobseekerDetails,
	}, nil

}

func (as *AdminServer) GetRecruiters(ctx context.Context, Req *pb.GetRecruiterRequest) (*pb.GetRecruitersResponse, error) {

	page := Req.Page

	GetRecruiters, err := as.adminUseCase.GetRecruiters(int(page))

	if err != nil {
		return &pb.GetRecruitersResponse{}, err
	}

	var recruitersDetails []*pb.RecruiterDetails
	for _, data := range GetRecruiters {
		recruitersDetails = append(recruitersDetails, &pb.RecruiterDetails{
			Id:          uint64(data.Id),
			Email:       data.Contact_mail,
			CompanyName: data.CompanyName,
			PhoneNumber: data.Phone,
			Blocked:     data.Blocked,
		})
	}
	return &pb.GetRecruitersResponse{
		Recruiter: recruitersDetails,
	}, nil

}

func (as *AdminServer) BlockRecruiter(ctx context.Context, Req *pb.BlockRecruiterRequest) (*pb.BlockRecruiterResponse, error) {

	recruiterId := Req.GetRecruiterId

	result, err := as.adminUseCase.BlockRecruiter(int(recruiterId))
	if err != nil {
		return &pb.BlockRecruiterResponse{}, err
	}
	return &pb.BlockRecruiterResponse{
		Status: result.Status,
	}, nil
}

func (as *AdminServer) UnBlockRecruiter(ctx context.Context, Req *pb.UnBlockRecruiterRequest) (*pb.UnBlockRecruiterResponse, error) {
	recruiterId := Req.GetRecruiterId

	result, err := as.adminUseCase.UnBlockRecruiter(int(recruiterId))
	if err != nil {
		return &pb.UnBlockRecruiterResponse{}, err
	}
	return &pb.UnBlockRecruiterResponse{
		Status: result.Status,
	}, nil
}

func (as *AdminServer) UnBlockJobseeker(ctx context.Context, Req *pb.UnBlockJobseekerRequest) (*pb.UnBlockJobseekerResponse, error) {
	jobseekerId := Req.JobseekerId

	result, err := as.adminUseCase.UnBlockJobseeker(int(jobseekerId))
	if err != nil {
		return &pb.UnBlockJobseekerResponse{}, err
	}
	return &pb.UnBlockJobseekerResponse{
		Status: result.Status,
	}, nil
}

func (as *AdminServer) BlockJobseeker(ctx context.Context, Req *pb.BlockJobseekerRequest) (*pb.BlockJobseekerResponse, error) {
	jobseekerId := Req.JobseekerId

	result, err := as.adminUseCase.BlockJobseeker(int(jobseekerId))
	if err != nil {
		return &pb.BlockJobseekerResponse{}, err
	}
	return &pb.BlockJobseekerResponse{
		Status: result.Status,
	}, nil
}

func (as *AdminServer) GetJobseekerDetails(ctx context.Context, Req *pb.GetJobseekerDetailsRequest) (*pb.GetJobseekerDetailsResponse, error) {
	jobseekerId := Req.JobseekerId

	result, err := as.adminUseCase.GetJobseekerDetails(int(jobseekerId))
	if err != nil {
		return &pb.GetJobseekerDetailsResponse{}, err
	}
	return &pb.GetJobseekerDetailsResponse{
		Id:          uint64(result.Id),
		Email:       result.Email,
		Firstname:   result.Name,
		PhoneNumber: result.Phone,
		Blocked:     result.Blocked,
	}, nil
}

func (as *AdminServer) GetRecruiterDetails(ctx context.Context, Req *pb.GetRecruiterDetailsRequest) (*pb.GetRecruiterDetailsResponse, error) {
	recruiterId := Req.RecruiterId

	result, err := as.adminUseCase.GetRecruiterDetails(int(recruiterId))
	if err != nil {
		return &pb.GetRecruiterDetailsResponse{}, err
	}
	return &pb.GetRecruiterDetailsResponse{
		Id:          uint64(result.Id),
		Email:       result.Contact_mail,
		CompanyName: result.CompanyName,
		PhoneNumber: result.Phone,
		Blocked:     result.Blocked,
	}, nil
}
